"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import io from "socket.io-client";
import Form from 'next/form';
import {orderAction, RoomData} from "./actions";

// import RoomDataa from "@/component/roomdata";

export default function RoomID() {
  const { id: roomId } = useParams();
  const [orderDay, setOrderDay] = useState(getTodayDate());
  const [checkin, setCheckIn] = useState("");
  const [checkout, setCheckOut] = useState("");
  const [messages, setMessages] = useState<string[]>([]);
  const [socket, setSocket] = useState<SocketIOClient.Socket | null>(null);

  // console.log(roomId);

  function getTodayDate() {
    const today = new Date();
    return today.toISOString().split("T")[0];
  }

  useEffect(() => {
    const socket = io("http://localhost:3000", { reconnection: true });
    setSocket(socket);

    socket.on("connect", () => socket.emit("joinRoom", roomId));
    socket.on("message", (msg: string) => setMessages((prev) => [...prev, msg]));
    return () => {
      socket.disconnect();
    };
  }, [roomId]);

  const sendRequest = (e: React.FormEvent) => {
    e.preventDefault();
    if (socket?.connected && checkin && checkout) {
      const requestMessage = `Order Day: ${orderDay}, Check-in: ${checkin}, Check-out: ${checkout}`;
      socket.emit("message", { roomId, message: requestMessage });
      setCheckIn("");
      setCheckOut("");
    }
  };

  const [roomname, setRoomName] = useState("");


  useEffect(()=>{
    const data = RoomData(roomId);
    data.then((v) => {
      setRoomName(JSON.parse(atob(v.room))) 
    })
  }, []);  

  return (
    <>
      {/* <RoomDataa roomId = {roomId}/> */}
      {/* <div className="Roomdata">
        {roomname.name}
      </div> */}

      <Form
        action={orderAction}
        // onSubmit={sendRequest}
        className="grid grid-cols-[30%_70%] w-96 gap-2 bg-red-300"
      >
        <input type="hidden" name="roomid" value={roomId} />
        <label htmlFor="date">date:</label>
        <input className="rounded-md border border-black" type="text" id="date" name="date" />
        <label htmlFor="from">from:</label>
        <input className="rounded-md border border-black" type="text" id="from" name="from" />
        <label htmlFor="to">to:</label>
        <input className="rounded-md border border-black" type="text" id="to" name="to" />
        <label htmlFor="state">state:</label>
        <input className="rounded-md border border-black" type="text" id="state" name="state" />
        <label htmlFor="note">note:</label>
        <input className="rounded-md border border-black" type="text" id="note" name="note" />
        <button type="submit" className="transition col-span-2 place-self-center w-36 h-8 rounded-md font-bold text-white bg-black hover:text-black hover:bg-white hover:ring-2 hover:ring-black duration-150">Dawtj</button>
      </Form>

      <div>
        <h2 className="text-xl font-bold mb-2">Room Request</h2>
        <ul className="border p-4 h-32 overflow-y-auto">
          {messages.map((msg, index) => (
            <li key={index} className="mb-2">{msg}</li>
          ))}
        </ul>
      </div>
    </>

  );
}