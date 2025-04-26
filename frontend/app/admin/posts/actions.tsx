'use server';

export async function GetPosts() {

    const res = await fetch(process.env.BACKEND_URL! + "/api/post/", {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
    });
    
    return await res.json();
}