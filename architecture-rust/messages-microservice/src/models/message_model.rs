use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct Message {
    #[serde(rename = "user1ID")]
    pub user_1_id: String,
    #[serde(rename = "user2ID")]
    pub user_2_id: String,
    #[serde(rename = "message")]
    pub message: String,
    #[serde(rename = "id")]
    pub id: Option<String>,
}
