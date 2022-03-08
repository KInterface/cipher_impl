export interface Encoded {
  Nonce: ArrayBuffer;
  Encrypted: ArrayBuffer;
  SharedKey: ArrayBuffer;
  Hash: string;
}
