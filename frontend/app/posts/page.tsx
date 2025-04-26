'use server';

import { ReadonlyEditor } from "@/component/ReadonlyEditor";

export default async function Posts() {
    const response = await fetch(process.env.BACKEND_URL! + "/api/post/", {
        method: "GET",
        // cache: "no-store",
    });

    const data = await response.json();

    console.log(data);

    const posts = JSON.parse(atob(data.posts));

    // console.log("POST: " + posts);
    return (
        <div className="flex flex-col gap-4 py-8">
            {posts && posts.map((post: { id: string, content: string }) => (
                <div id={post.id}>
                    {/* <ReadonlyEditor data={post.content} /> */}
                    {post.content}
                </div>
            ))}
        </div>
    );
}