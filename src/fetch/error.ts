import axios from "axios";

export const handleError = <T>(e: unknown): FetchError<T> => {
  if (axios.isAxiosError(e)) {
    return e.response
      ? new FetchError(e.response.status, e.message, e.response.data)
      : new FetchError(500, e.message);
  } else {
    return new FetchError(500, (e as Error).message);
  }
};

class FetchError<T> extends Error {
  constructor(
    readonly status: number,
    readonly message: string,
    readonly data?: T
  ) {
    super();
  }
}
