import { redirect } from "next/dist/server/api-utils";
import ExampleImg from "../../public/exampleimg/room_img.jpg"
import Image from "next/image";



export default async function Rooms() {
  const data = await fetch(process.env.BACKEND_URL! + "/api/room/")
  return (
    <div className="container flex flex-col h-full mx-auto gap-8">
      {/* {data.text()} */}

      <div className="lg:w-[70%] w-full h-auto min-h-96 mx-auto md:grid md:grid-cols-[60%_40%] flex flex-col gap-1 rounded-md shadow-lg">
        <Image src={ExampleImg} alt="logo" className="rounded-tl-md" />
        <a href="/rooms/test_id" className="translate duration-300 text-3xl font-bold text-center content-center ">
          <p className="hover:text-green-800">
            ROOM NAME
          </p>
        </a>
        <div className="mx-[5px] min-h-12 text-lg text-center content-center truncate text-wrap text-ellipsis">
          403 QL13, Hiệp Bình Phước, Thủ Đức, Hồ Chí Minh, Việt Nam
        </div>
        <div className="mx-[5px] min-h-12 text-lg text-center content-center truncate text-wrap text-ellipsis">
          DESCRIPTION
        </div>
        <div className="flex justify-center items-center md:order-none order-last md:h-24 h-12 m-[5px]">
          <a className="translate duration-300 h-12 w-40 rounded-lg text-lg text-center content-center bg-black text-white hover:text-black hover:bg-white hover:ring-2 hover:ring-black" href="/rooms/test_id">
            Đặt phòng
          </a>
        </div>
        <div className="mx-[5px] min-h-12 text-lg font-bold text-center content-center truncate text-wrap text-ellipsis">
          120k/h
        </div>
      </div>


    </div>
  );
}

