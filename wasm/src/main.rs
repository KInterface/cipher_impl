use anyhow::{bail, Result};
use argon2::{
    password_hash::{rand_core::OsRng, Output, PasswordHasher, SaltString},
    Argon2,
};

fn main(){
   let a=vec![4, 79, 146, 243, 235, 215, 103, 10, 92, 173, 127, 193, 153, 41, 218, 147, 162, 68, 164, 44, 42, 69, 40, 36, 24, 90, 227, 172, 92, 183, 76, 134, 111, 237, 169, 229, 114, 46, 119, 39, 201, 188, 144, 242, 242, 166, 250, 137, 147, 160, 26, 24, 36, 216, 228, 141, 29, 159, 135, 238, 22, 179, 224, 182, 166];
 let (shared_secret, pub_key) = calcurate_shared_secret(&a);


}




pub fn calcurate_shared_secret(pk: &[u8]) -> (SharedSecret, PublicKey) {
    let pk_bytes = EncodedPoint::from_bytes(pk).expect("invalid");
    let received_pub_key =
        PublicKey::from_sec1_bytes(pk_bytes.as_ref()).expect("this public key is invalid!");
    let disposable_secret = EphemeralSecret::random(&mut OsRng);
    let shared_key = disposable_secret.public_key();
    (
        disposable_secret.diffie_hellman(&received_pub_key),
        shared_key,
    )
}
