use anyhow::{anyhow, Result};
use p256::{
    ecdh::{EphemeralSecret, SharedSecret},
    EncodedPoint, PublicKey,
};
use rand_core::OsRng;

pub fn calcurate_shared_secret(pk: &[u8]) -> Result<(SharedSecret, PublicKey)> {
    let pk_bytes = EncodedPoint::from_bytes(pk)
        .map_err(|e| anyhow!("failed generate encoded point from pub_key: {}", e))?;
    let received_pub_key = PublicKey::from_sec1_bytes(pk_bytes.as_ref())?;
    let disposable_secret = EphemeralSecret::random(&mut OsRng);
    let shared_key = disposable_secret.public_key();
    Ok((
        disposable_secret.diffie_hellman(&received_pub_key),
        shared_key,
    ))
}
