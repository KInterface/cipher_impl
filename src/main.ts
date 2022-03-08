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
  console.log(encoded);
  const res = await post<TokenResponse, any, { Data: string } & Encoded>({
    Data: card,
    ...encoded,
  });
  if (res) {
    const tokenInput = document.getElementById("token") as HTMLInputElement;
    tokenInput.value = res.Token;
  }
};

window.onload = async () => {
  const submit = document.getElementById("submit");
  submit && submit.addEventListener("click", fetchToken);
  const pubKey = await get<ArrayBuffer, any>("", "arraybuffer");
  if (pubKey) {
    console.log(pubKey);
    console.log(new Uint8Array(pubKey));
    init().then(() => {
      console.log(encode("", new Uint8Array(pubKey)));
    });
    PublicKey = pubKey;
  }
};
