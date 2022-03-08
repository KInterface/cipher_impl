import axios, { AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";
import { ENDPOINT_URL } from "../constants/endpoint";
import { handleError } from "./error";

export const get = async <T, D>(
  path: string = "",
  responseType: ResponseType = "json"
): Promise<T | void> => {
  const options: AxiosRequestConfig = {
    url: ENDPOINT_URL + path,
    method: "GET",
    responseType,
  };
  try {
    const res: AxiosResponse<T, D> = await axios(options);
    return res.data;
  } catch (e) {
    handleError(e);
  }
};

export const post = async <T, D, U>(
  body: U,
  path: string = "",
  responseType: ResponseType = "json"
): Promise<T | void> => {
  const options: AxiosRequestConfig = {
    url: ENDPOINT_URL + path,
    method: "POST",
    responseType,
    data: body,
  };
  try {
    const res: AxiosResponse<T, D> = await axios(options);
    return res.data;
  } catch (e) {
    handleError<T>(e);
  }
};
