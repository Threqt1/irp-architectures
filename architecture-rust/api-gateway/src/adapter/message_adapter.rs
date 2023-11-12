use reqwest::{Client, StatusCode};
use rocket::{http::Status, response::status};

pub struct MessageServiceAdapter {
    pub client: Client,
    pub route: String,
}

impl MessageServiceAdapter {
    pub async fn send_message(
        &self,
        user1: String,
        user2: String,
        content: String,
    ) -> Result<crate::models::message_model::Message, status::Custom<String>> {
        let local_client = self.client.clone();

        let message = crate::models::message_model::Message {
            user_1_id: user1,
            user_2_id: user2,
            message: content,
            id: String::new(),
        };

        let message_response = local_client
            .put(self.route.as_str())
            .json(&message)
            .send()
            .await;

        if let Err(_) = message_response {
            return Err(status::Custom(
                Status::InternalServerError,
                String::from("message: internal message server error"),
            ));
        }

        let message_response = message_response.unwrap();

        if message_response.status() != StatusCode::OK {
            return Err(status::Custom(
                Status::InternalServerError,
                String::from("message: internal message server response"),
            ));
        }

        let message_response = message_response
            .json::<crate::models::message_model::Message>()
            .await;

        if let Err(_) = message_response {
            return Err(status::Custom(
                Status::InternalServerError,
                String::from("message: internal message JSON"),
            ));
        }

        Ok(message_response.unwrap())
    }
}
