// "use client"

import ExampleImg from "../../public/exampleimg/room_img.jpg"
import Image from "next/image";

interface timestamp {
  Time: string,
  Status: number,
  InfinityModifier: number
}

interface Room {
  id: string,
  name: string,
  address: string,
  description: string,
  price: number,
  createdate: timestamp,
  editdate: timestamp,
  userid: number
}

export default async function Rooms() {
  const data = await fetch(process.env.BACKEND_URL! + "/api/room/")
  const response: Room[] = JSON.parse(atob((await data.json()).rooms))
  return (
    <div className="container flex flex-col h-full mx-auto gap-8 py-10">
      {response.map((val) =>
        <div key={val.id} className="lg:w-[70%] w-full h-auto min-h-96 mx-auto md:grid md:grid-cols-[60%_40%] flex flex-col gap-1 rounded-md shadow-lg">
          <Image src={ExampleImg} alt="logo" className="rounded-tl-md" />
          <a href={`/rooms/` + val.id} className="translate duration-300 text-3xl font-bold text-center content-center ">
            <p className="hover:text-green-800">
              {val.name}
            </p>
          </a>
          <div className="mx-[5px] min-h-12 text-lg text-center content-center truncate text-wrap text-ellipsis">
            {val.address}
          </div>
          <div className="mx-[5px] min-h-12 text-lg text-center content-center truncate text-wrap text-ellipsis">
            {val.description}
          </div>
          <div className="flex justify-center items-center md:order-none order-last md:h-24 h-12 m-[5px]">
            <a className="translate duration-300 h-12 w-40 rounded-lg text-lg text-center content-center bg-black text-white hover:text-black hover:bg-white hover:ring-2 hover:ring-black" href={`/rooms/${val.id}`}>
              Đặt phòng
            </a>
          </div>
          <div className="mx-[5px] min-h-12 text-lg font-bold text-center content-center truncate text-wrap text-ellipsis">
            {val.price} k/h
          </div>
        </div>
      )}


    </div>
  );
}

