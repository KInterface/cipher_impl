use anyhow::{bail, Result};
use argon2::{
    password_hash::{rand_core::OsRng, Output, PasswordHasher, SaltString},
    Argon2,
};

pub fn generate(password: &[u8]) -> Result<(Output, String)> {
    let salt = SaltString::generate(&mut OsRng);
    let argon2 = Argon2::default();
    let hash_result = argon2.hash_password(password, &salt);
    let hash = match hash_result {
        Ok(h) if h.hash.is_some() => h,
        Err(e) => bail!("error occured while generating hash {}", e),
        _ => bail!("Something went wrong."),
    };
    if let Some(h) = hash.hash {
        Ok((h, hash.to_string()))
    } else {
        bail!("hashed result is empty!")
    }
}
