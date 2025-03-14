"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import io from "socket.io-client";
import {Form} from ""

export default function RoomID() {
  const { id: roomId } = useParams();
  const [orderDay, setOrderDay] = useState(getTodayDate());
  const [checkin, setCheckIn] = useState("");
  const [checkout, setCheckOut] = useState("");
  const [messages, setMessages] = useState<string[]>([]);
  const [socket, setSocket] = useState<SocketIOClient.Socket | null>(null);

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

  // So sánh giờ checkin và checkout
  const isValidCheckOut = (checkinTime: string, checkoutTime: string) =>
    checkinTime && checkoutTime && hours.indexOf(checkinTime) < hours.indexOf(checkoutTime);

  return (
    <>
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