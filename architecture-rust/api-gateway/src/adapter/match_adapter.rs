use reqwest::{Client, StatusCode};
use rocket::{http::Status, response::status};

pub struct MatchServiceAdapter {
    pub client: Client,
    pub route: String,
}

impl MatchServiceAdapter {
    pub async fn match_user(
        &self,
        id: &str,
    ) -> Result<crate::models::match_model::Match, status::Custom<String>> {
        let local_client = self.client.clone();

        let match_response = local_client
            .put(self.route.as_str())
            .query(&[("id", id)])
            .send()
            .await;

        if let Err(_) = match_response {
            return Err(status::Custom(
                Status::InternalServerError,
                String::from("match_user: internal match server error"),
            ));
        }

        let match_response = match_response.unwrap();

        if match_response.status() != StatusCode::OK {
            return Err(status::Custom(
                Status::InternalServerError,
                String::from("match_user: internal match server response"),
            ));
        }

        let match_response = match_response
            .json::<crate::models::match_model::Match>()
            .await;

        if let Err(_) = match_response {
            return Err(status::Custom(
                Status::InternalServerError,
                String::from("match_user: internal match JSON error"),
            ));
        }

        Ok(match_response.unwrap())
    }

    pub async fn get_match(
        &self,
        id: &str,
    ) -> Result<crate::models::match_model::Match, status::Custom<String>> {
        let local_client = self.client.clone();

        let match_response = local_client
            .get(self.route.as_str())
            .query(&[("id", id)])
            .send()
            .await;

        if let Err(_) = match_response {
            return Err(status::Custom(
                Status::InternalServerError,
                String::from("get_match: internal match server error"),
            ));
        }

        let match_response = match_response.unwrap();

        if match_response.status() != StatusCode::OK {
            return Err(status::Custom(
                Status::InternalServerError,
                String::from("get_match: internal match server response"),
            ));
        }

        let match_response = match_response
            .json::<crate::models::match_model::Match>()
            .await;

        if let Err(_) = match_response {
            return Err(status::Custom(
                Status::InternalServerError,
                String::from("get_match: internal match JSON"),
            ));
        }

        Ok(match_response.unwrap())
    }
}
