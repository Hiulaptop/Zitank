export default async function Home() {
    const data = await fetch(process.env.BACKEND_URL!)
    return (
      <div className="">
          {data.text()}
      </div>
    );
  }
  