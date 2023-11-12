use std::{
    collections::{HashMap, VecDeque},
    sync::{Mutex, RwLock},
};

use rocket::{get, http::Status, launch, put, response::status, routes, serde::json::Json, State};

mod models;

const DEFAULT_RETRY_TIME: u64 = 10000000;

struct APIConfig {
    queue: Mutex<VecDeque<String>>,
    matches: RwLock<HashMap<String, String>>,
}

#[put("/match?<id>")]
fn match_user(
    config: &State<APIConfig>,
    id: String,
) -> Result<status::Custom<Json<crate::models::match_model::Match>>, status::Custom<String>> {
    let mut queue = config.queue.lock().unwrap();
    let matched = queue.pop_front();
    match matched {
        Some(matched_id) => {
            let mut map = config.matches.write().unwrap();
            map.insert(matched_id.clone(), id);
            Ok(status::Custom(
                Status::Ok,
                Json(crate::models::match_model::Match {
                    matched_user_id: matched_id,
                    wait_for: 0,
                }),
            ))
        }
        None => {
            queue.push_back(id);
            Ok(status::Custom(
                Status::Ok,
                Json(crate::models::match_model::Match {
                    matched_user_id: String::from(""),
                    wait_for: DEFAULT_RETRY_TIME,
                }),
            ))
        }
    }
}

#[get("/match?<id>")]
fn get_user(
    config: &State<APIConfig>,
    id: String,
) -> Result<status::Custom<Json<crate::models::match_model::Match>>, status::Custom<String>> {
    let map = config.matches.read().unwrap();
    let matched = map.get(id.as_str());
    match matched {
        Some(matched_id) => Ok(status::Custom(
            Status::Ok,
            Json(crate::models::match_model::Match {
                matched_user_id: String::from(matched_id),
                wait_for: 0,
            }),
        )),
        None => Ok(status::Custom(
            Status::Ok,
            Json(crate::models::match_model::Match {
                matched_user_id: String::from(""),
                wait_for: DEFAULT_RETRY_TIME,
            }),
        )),
    }
}

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![match_user, get_user])
        .manage(APIConfig {
            queue: Mutex::new(VecDeque::from(vec![String::from("0")])),
            matches: RwLock::new(HashMap::new()),
        })
}
