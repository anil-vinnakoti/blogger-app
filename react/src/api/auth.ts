import { useMutation } from "@tanstack/react-query";
import { httpRequest } from "./client";

interface User {
  id: number;
  username: string;
  name: string;
}

interface LoginCredentials {
  username: string;
  password: string;
}

export function useLoginOrSignUp(url: string) {
  return useMutation<User, string, LoginCredentials>({
    mutationFn: async (credentials: LoginCredentials) => {
      const { data } = await httpRequest({
        method: "post",
        url,
        data: credentials
      })
        .then((res) => Promise.resolve(res.data))
        .catch((err) => Promise.reject(err.data.error));
      return data; // must match User
    }
  });
}

// Logout mutation
export function useLogout() {
  return useMutation<void, string, void>({
    mutationFn: async (): Promise<void> => {
      await httpRequest({ method: "post", url: "/logout" })
        .then((res) => Promise.resolve(res.data))
        .catch((err) => Promise.reject(err.data.error));
    }
  });
}

// Session check (just async function, not a mutation)
export async function checkSession(): Promise<User> {
  const { data } = await httpRequest({ method: "get", url: "/me" });
  return data;
}
