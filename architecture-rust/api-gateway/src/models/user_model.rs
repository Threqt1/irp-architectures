use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct User {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "match")]
    pub matched_user: Option<String>,
}
