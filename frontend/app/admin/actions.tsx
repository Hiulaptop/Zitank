"use server";

import { cookies } from "next/headers";

export async function AdminData(){
    const cookieStore = await cookies();

    const res = await fetch(process.env.BACKEND_URL! + "/api/user/admin/",{
        method: "GET",
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `BEARER ${cookieStore.get("jwt")!.value}`
        },
    });

    return await res.json();
}