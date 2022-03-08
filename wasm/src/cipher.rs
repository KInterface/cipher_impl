mod hash;
mod secp256r1;
mod xchacha20;
use anyhow::Result;
use argon2::password_hash::Output;
use p256::PublicKey;

pub fn enc_xchacha20(msg: &str, key: &[u8; 32]) -> (Vec<u8>, Vec<u8>) {
    xchacha20::enc(msg, key)
}

pub fn enc_secp256r1(pk_bytes_slice: &[u8]) -> Result<(Output, String, PublicKey)> {
    let (shared_secret, pub_key) = secp256r1::calcurate_shared_secret(&pk_bytes_slice);
    let (private_key, hash) = hash::generate(shared_secret.as_bytes())?;
    Ok((private_key, hash, pub_key))
}
