use models::message_model::Message;
use rocket::{http::Status, launch, put, response::status, routes, serde::json::Json, State};

mod library;
mod models;

struct APIConfig {
    snowflake_provider: crate::library::snowflake::SnowflakeProvider,
}

#[put("/message", data = "<message>")]
fn create_message(
    config: &State<APIConfig>,
    message: Json<Message>,
) -> Result<status::Custom<Json<Message>>, status::Custom<String>> {
    let id = config.snowflake_provider.generate();
    let parsed_message = message.into_inner();
    let message = Message {
        user_1_id: parsed_message.user_1_id,
        user_2_id: parsed_message.user_2_id,
        message: parsed_message.message,
        id: Some(id),
    };

    Ok(status::Custom(Status::Ok, Json(message)))
}

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![create_message])
        .manage(APIConfig {
            snowflake_provider: crate::library::snowflake::SnowflakeProvider {
                generator: snowdon::Generator::default(),
            },
        })
}
