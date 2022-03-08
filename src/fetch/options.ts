import { AxiosRequestHeaders } from "axios";

export type WithTokenHeader = AxiosRequestHeaders & Readonly<{ token: string }>;
