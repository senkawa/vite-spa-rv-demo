import { useEffect, useState } from "react";
import { Post, Response } from "@/typings/response";
import axios from "axios";

export default function Notes() {
  const [posts, setPost] = useState<Post[]>([]);

  const fetchPosts = async () => {
    let { data } = await axios.get<Response>("http://localhost:8000");
    setPost(data.data);
  };

  useEffect(() => {
    fetchPosts();
  }, []);

  if (posts.length > 0) {
    return (
      <main className="p-1">
        <h2 className="font-medium">Notes Made</h2>
        <ul>
          {posts.map((post) => (
            <li key={post.ID}>
              #{post.ID} {post.Body}
            </li>
          ))}
        </ul>
      </main>
    );
  }

  return (
    <main>
      <p>No posts</p>
    </main>
  );
}
