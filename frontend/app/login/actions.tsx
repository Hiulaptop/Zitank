'use server';

import { redirect } from 'next/navigation';

export async function loginAction(formData: FormData) {
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
}
