mod cipher;
pub mod js;
use elliptic_curve::sec1::ToEncodedPoint;
use js_sys::{JsString, Uint8Array};
use wasm_bindgen::prelude::*;

#[wasm_bindgen]
pub fn encode(message: JsString, public_key: Uint8Array) -> JsValue {
    let pk_vec = js::array_to_vec(public_key);
    let (private, hash, public) = match cipher::enc_secp256r1(&pk_vec[..]){
       Ok((output,hash,pub_key))=>(output,hash,pub_key) ,
       Err(e)=>panic!("{}", format!("something went wrong during generate shared private: {}",e))
    };
    let pk_shared_point = public
        .as_affine()
        .to_encoded_point(false)
        .as_bytes()
        .to_vec();
    let b_secret = private.as_bytes();
    let adjusted_secret: [u8; 32] = b_secret.try_into().unwrap_or([0x8;32]);

    let msg: String = message.into();
    let (nonce, encrypted) = cipher::enc_xchacha20(&msg, &adjusted_secret);
    let encrypted = js::Encrypted {
        Nonce: nonce,
        Encrypted: encrypted,
        SharedKey: pk_shared_point,
        Hash: hash,
    };

    JsValue::from_serde(&encrypted).unwrap()
}
