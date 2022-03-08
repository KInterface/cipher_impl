use p256::{
    ecdh::{EphemeralSecret, SharedSecret},
    EncodedPoint, PublicKey,
};
use rand_core::OsRng;

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
