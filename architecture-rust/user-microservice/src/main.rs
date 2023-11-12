use ::rocket::launch;
use rocket::{http::Status, put, response::status, routes, serde::json::Json, State};

mod library;
mod models;

struct APIConfig {
    snowflake_provider: crate::library::snowflake::SnowflakeProvider,
}

#[put("/users")]
fn create_user(
    config: &State<APIConfig>,
) -> Result<status::Custom<Json<crate::models::user_model::User>>, status::Custom<String>> {
    let id = config.snowflake_provider.generate();
    let user = crate::models::user_model::User { id: id };

    Ok(status::Custom(Status::Ok, Json(user)))
}

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![create_user])
        .manage(APIConfig {
            snowflake_provider: crate::library::snowflake::SnowflakeProvider {
                generator: snowdon::Generator::default(),
            },
        })
}
