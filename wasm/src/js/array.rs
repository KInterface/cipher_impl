use js_sys::{Array, Uint8Array};
use wasm_bindgen::JsValue;

pub fn vec_to_array(vec: Vec<u8>) -> Uint8Array {
    let js_array: Array = vec.into_iter().map(JsValue::from).collect();
    Uint8Array::new(&js_array)
}

pub fn array_to_vec(array: Uint8Array) -> Vec<u8> {
    array.to_vec()
}
