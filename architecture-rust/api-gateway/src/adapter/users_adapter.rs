use reqwest::{Client, StatusCode};
use rocket::{http::Status, response::status};

pub struct UsersServiceAdapter {
    pub client: Client,
    pub route: String,
}

impl UsersServiceAdapter {
    pub async fn create_user(
        &self,
    ) -> Result<crate::models::user_model::User, status::Custom<String>> {
        let local_client = self.client.clone();

        let user_response = local_client.put(self.route.as_str()).send().await;

        if let Err(_) = user_response {
            return Err(status::Custom(
                Status::InternalServerError,
                String::from("user: internal user server error"),
            ));
        }

        let user_response = user_response.unwrap();

        if user_response.status() != StatusCode::OK {
            return Err(status::Custom(
                Status::InternalServerError,
                String::from("user: internal user server response"),
            ));
        }

        let user_response = user_response
            .json::<crate::models::user_model::User>()
            .await;

        if let Err(e) = user_response {
            return Err(status::Custom(
                Status::InternalServerError,
                String::from(e.to_string()),
            ));
        }

        Ok(user_response.unwrap())
    }
}
