import { useMutation } from "@tanstack/react-query";
import api from "./client";

interface User {
  id: number;
  username: string;
  name: string;
}

interface LoginCredentials {
  username: string;
  password: string;
}

export function useLogin() {
  return useMutation<User, Error, LoginCredentials>({
    mutationFn: async (credentials: LoginCredentials) => {
      const { data } = await api.post<User>("/login", credentials);
      return data; // must match User
    }
  });
}

// Logout mutation
export function useLogout() {
  return useMutation<void, Error, void>({
    mutationFn: async (): Promise<void> => {
      await api.post("/logout");
    }
  });
}

// Session check (just async function, not a mutation)
export async function checkSession(): Promise<User> {
  const { data } = await api.get<User>("/me");
  return data;
}
