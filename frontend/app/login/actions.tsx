'use server';

import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

export async function loginAction(formData: FormData) {
  const cookieStore = await cookies()
  const username = formData.get('username');
  const password = formData.get('password');

  const response = await fetch(process.env.BACKEND_URL! + "/api/user/login", {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ username, password }),
  });

  if (!response.ok) {
    redirect("/error");
  }
  cookieStore.set("jwt", (await response.json()).token)
  redirect('/')
}
