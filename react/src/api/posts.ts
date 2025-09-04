import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import api from "./client";

export function usePosts() {
  return useQuery(["posts"], async () => {
    const { data } = await api.get("/posts");
    return data;
  });
}

export function useCreatePost() {
  const queryClient = useQueryClient();

  return useMutation(
    async (newPost: { title: string; content: string }) => {
      const { data } = await api.post("/posts", newPost);
      return data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries(["posts"]);
      }
    }
  );
}
