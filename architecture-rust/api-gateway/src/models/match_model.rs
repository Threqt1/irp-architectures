use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct Match {
    #[serde(rename = "matchedUserID")]
    pub matched_user_id: String,
    #[serde(rename = "waitFor")]
    pub wait_for: u64,
}
