// import { usePosts, useCreatePost } from "../api/posts";

export default function Posts() {
  // const { data: posts, isLoading } = usePosts();
  // const createPost = useCreatePost();

  // if (isLoading) return <p>Loading...</p>;

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">Posts</h1>
      {/* <button
        onClick={() =>
          createPost.mutate({ title: "New Post", content: "Hello World!" })
        }
        className="bg-green-600 text-white px-4 py-2 rounded"
      >
        Add Post
      </button>

      <ul className="mt-4 space-y-2">
        {posts?.map((p: any) => (
          <li key={p.id} className="p-4 border rounded">
            <h2 className="font-semibold">{p.title}</h2>
            <p>{p.content}</p>
          </li>
        ))}
      </ul> */}
    </div>
  );
}
