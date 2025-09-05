import axios, { AxiosRequestConfig, AxiosResponse } from "axios";

export class ServerError extends Error {
  message = "";
  name = "ServerError";
  info = {};
  constructor(message: string, info: any) {
    super(message);
    this.message = message;
    this.info = info;
  }
}

export const axiosInstance = axios.create({
  baseURL: "http://localhost:8080/api",
  withCredentials: true, // session cookie handling
  timeout: 10000
});

export async function httpRequest<T = any>(
  options: AxiosRequestConfig
): Promise<AxiosResponse<T>> {
  try {
    const response = await axiosInstance(options);

    return response;
  } catch (error: any) {
    if (error.response) {
      console.error("❌ Request Failed:", error.response);

      if (error.response.status >= 500) {
        throw new ServerError("Status: 500 | API Error", error.response);
      }

      // always reject with response for consistency
      return Promise.reject(error.response);
    }

    console.error("Error Message:", error.message);
    return Promise.reject({
      data: { errors: [{ errorMessage: error.message }] }
    });
  }
}
