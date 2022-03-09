import { get, post } from "./fetch/fetcher";
import { Card } from "./types/card";
import { TokenResponse } from "./types/token";
import init, { encode } from "../wasm/pkg";
import { Encoded } from "./types/encoded";

let PublicKey: ArrayBuffer;
export const fetchToken = async () => {
  const no = (document.getElementById("no") as HTMLInputElement)?.value ?? "";
  const name =
    (document.getElementById("name") as HTMLInputElement)?.value ?? "";
  const expiryDate =
    (document.getElementById("expiry_date") as HTMLInputElement)?.value ?? "";
  const code =
    (document.getElementById("code") as HTMLInputElement)?.value ?? "";

  const cardObj: Card = {
    no,
    name,
    expiryDate,
    code,
  };

  const card = Object.entries(cardObj)
    .map(([key, value]) => `${key}:${value}`)
    .join("|");
  const encoded = encode(card, new Uint8Array(PublicKey)) as Encoded;
  const res = await post<TokenResponse, any, { Data: string } & Encoded>({
    Data: card,
    ...encoded,
  });
  if (res) {
    const tokenInput = document.getElementById("token") as HTMLInputElement;
    tokenInput.value = res.Token;
    const resInput = document.getElementById("res") as HTMLInputElement;
    resInput.innerHTML += `<p>Server received following encrypted data: ${res.Encoded}</p>`;
    resInput.innerHTML += `<p>Server decrypt and get following data: ${res.Decoded}</p>`;
    resInput.innerHTML += `<p>Token: ${res.Token}</p>`;
  }
};

window.onload = async () => {
  const submit = document.getElementById("submit");
  submit && submit.addEventListener("click", fetchToken);
  const pubKey = await get<ArrayBuffer, any>("", "arraybuffer");
  if (pubKey) {
    init();
    PublicKey = pubKey;
  }
};
