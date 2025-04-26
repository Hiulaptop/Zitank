'use server';
import { createHmac } from 'crypto';
import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';

export async function orderAction(formData: FormData) {
    const cookieStore = await cookies()
    const roomid = formData.get("roomid");
    const from = formData.get("from");
    const to = formData.get("to");
    const state = formData.get("state");
    const note = formData.get("note");

    const res = await fetch(process.env.BACKEND_URL! + `/api/room/${roomid}/order`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `BEARER ${cookieStore.get("jwt")!.value}`
        },
        body: JSON.stringify({
            'fromto': `[${from},${to}]`,
            'state': state,
            'note': note
        }),
    });
    if (!res.ok) {
        console.log(await res.text())
        redirect("/error");
    }
    const hmac = createHmac("sha256", process.env.CHECKSUM_KEY!)
    let resp = await res.json()
    hmac.update(`amount=${resp.totalprice}&cancelUrl=${process.env.CANCEL_URL!}&description=TLTH12AG&orderCode=${resp.id}&returnUrl=${process.env.RETURN_URL!}`)
    const response = await fetch(`${process.env.PAYOS_API!}/v2/payment-requests`, {
        method: "POST",
        headers: {
            'x-client-id': process.env.CLIENT_ID!,
            'x-api-key': process.env.API_KEY!,
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "orderCode": resp.id,
            "amount": resp.totalprice,
            "description": "TLTH12AG",
            "cancelUrl": process.env.CANCEL_URL!,
            "returnUrl": process.env.RETURN_URL!,
            "expiredAt": Math.floor(Date.now() / 1000) + 60 * 10,
            "signature": hmac.digest('hex')
        })
    })
    resp = await response.json()
    if (resp.code != "00") {
        redirect("/err")
    }
    redirect(resp.data.checkoutUrl)
}

export async function RoomData(roomId: any){
    // const cookieStore = await cookies()
    const res = await fetch(process.env.BACKEND_URL! + `/api/room/${roomId}`,{
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
    });

    return await res.json()
}