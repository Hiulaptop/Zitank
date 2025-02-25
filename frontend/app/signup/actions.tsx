'use server';

import { redirect } from 'next/navigation';

export async function signupAction(formData: FormData) {
    const Username = formData.get('Username');
    const Password = formData.get('Password');
    const Fullname = formData.get('Fullname');
    const Email = formData.get('Email');
    const PhoneNumber = formData.get('PhoneNumber');
    const Gender = formData.get('Gender');
    const Role = "admin";

    const response = await fetch(process.env.BACKEND_URL! + "/api/user/register", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            Username,
            Password,
            Fullname,
            Email,
            PhoneNumber,
            Gender,
            Role,
        }),
    });

    if (!response.ok) {
        redirect("/error");
    }
    else{
        redirect('/login');
    }
}
