export type Token = string & { _token: never };

export interface TokenResponse {
  Token: Token;
  Encoded: string;
  Decoded: string;
}
