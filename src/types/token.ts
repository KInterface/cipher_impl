export type Token = string & { _token: never };

export interface TokenResponse {
  Token: Token;
}
