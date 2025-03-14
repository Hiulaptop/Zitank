'use server';
import { createHmac } from "crypto";
import { redirect } from 'next/navigation';

//Khong nen de o client
async function createPayment(amount: number, orderID: number) {
    const hmac = createHmac("sha256", process.env.CHECKSUM_KEY!)
    hmac.update(`amount=${amount}&cancelUrl=${process.env.CANCEL_URL!}&description=TLTH12AG&orderCode=${orderID}&returnUrl=${process.env.RETURN_URL!}`)
    const response = await fetch(`${process.env.PAYOS_API!}/v2/payment-requests`, {
        method: "POST",
        headers: {
            'x-client-id': process.env.CLIENT_ID!,
            'x-api-key': process.env.API_KEY!
        },
        body: JSON.stringify({
            "orderCode": orderID,
            "amount": amount,
            "description": "TLTH12AG",
            "cancelUrl": process.env.CANCEL_URL!,
            "returnUrl": process.env.RETURN_URL!,
            "expiredAt": Math.floor(Date.now() / 1000) + 60 * 10,
            "signature": hmac.digest('hex')
        })
    })
    let res = await response.json()
    if (res.code != "00")
        redirect("/err")
    redirect(res.data.checkoutUrl)
}

export default async function Payment(orderID: number) {
    const response = await fetch(`${process.env.BACKEND_URL}/api/order/${orderID}/`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `BEARER ${localStorage.getItem("jwt")}`
        }
    })
    if (response.status != 200)
        redirect("/err")
    let js = await response.json()

    return (
        <div>
            <button onClick={() => createPayment(js.totalprice, orderID)}>
                Test
            </button>
        </div>
    )
}