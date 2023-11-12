use rocket::http::Status;
use rocket::serde::json::Json;
use rocket::State;
use rocket::{launch, post, response::status, routes};
use tokio::time::{sleep, Duration};

mod adapter;
mod config;
mod models;

struct APIConfig {
    users_service_adapter: crate::adapter::users_adapter::UsersServiceAdapter,
    messages_service_adapter: crate::adapter::message_adapter::MessageServiceAdapter,
    match_service_adapter: crate::adapter::match_adapter::MatchServiceAdapter,
}

#[post("/test")]
async fn test(
    config: &State<APIConfig>,
) -> Result<status::Custom<Json<crate::models::message_model::Message>>, status::Custom<String>> {
    let user = config.users_service_adapter.create_user().await?;

    let mut matched = config
        .match_service_adapter
        .match_user(user.id.as_str())
        .await?;

    while matched.wait_for > 0 {
        sleep(Duration::from_nanos(matched.wait_for)).await;
        matched = config
            .match_service_adapter
            .get_match(user.id.as_str())
            .await?;
    }

    let message = config
        .messages_service_adapter
        .send_message(
            user.id,
            matched.matched_user_id,
            String::from("Hello Other User"),
        )
        .await?;

    Ok(status::Custom(Status::Ok, Json(message)))
}

#[launch]
fn rocket() -> _ {
    let client = reqwest::Client::builder().build().unwrap();

    rocket::build().mount("/", routes![test]).manage(APIConfig {
        users_service_adapter: crate::adapter::users_adapter::UsersServiceAdapter {
            client: client.clone(),
            route: String::from(crate::config::config::USERS_MS_ROUTE),
        },
        messages_service_adapter: crate::adapter::message_adapter::MessageServiceAdapter {
            client: client.clone(),
            route: String::from(crate::config::config::MESSAGING_MS_ROUTE),
        },
        match_service_adapter: crate::adapter::match_adapter::MatchServiceAdapter {
            client: client.clone(),
            route: String::from(crate::config::config::MATCHING_MS_ROUTE),
        },
    })
}
