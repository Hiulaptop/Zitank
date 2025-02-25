


export default async function Page({ params,}: { params: Promise<{ id: string }> }) {
    const id = (await params).id
    const data = await fetch(process.env.BACKEND_URL! +"/room/" + id)

    return <div className="text-black">{data.text()}</div>
}