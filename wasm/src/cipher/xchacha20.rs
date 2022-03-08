use aead::Aead;
use chacha20poly1305::aead::NewAead;
use chacha20poly1305::{Key, XChaCha20Poly1305, XNonce};
use rand_chacha::ChaCha20Rng;
use rand_core::{RngCore, SeedableRng};

// Receive secret key and msg you want to encrypt.
// Return a tuple contains (nonce,enrypted message)
pub fn enc(msg: &str, key: &[u8; 32]) -> (Vec<u8>, Vec<u8>) {
    // Convert args to bytes.
    let b_key = Key::from_slice(key);
    let b_msg = msg.as_bytes();

    // Generate new Xchacha20-Poly1305 instance.
    let aead = XChaCha20Poly1305::new(&b_key);

    // Generate random nonce.
    // Xchacha20-Poly1305 with random nonce causes secret-key wear-out after sending 2^80 messages.
    let mut data = [0u8; 24];
    data = gen_rand(data);
    let nonce = XNonce::from_slice(&data);

    (
        nonce.to_vec(),
        aead.encrypt(nonce, b_msg.as_ref()).unwrap_or(Vec::new()),
    )
}

fn gen_rand<const N: usize>(mut slice: [u8; N]) -> [u8; N] {
    ChaCha20Rng::from_entropy().fill_bytes(&mut slice);
    slice
}
