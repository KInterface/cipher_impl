use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct Encrypted {
    pub Nonce: Vec<u8>,
    pub Encrypted: Vec<u8>,
    pub SharedKey: Vec<u8>,
    pub Hash:String
}
