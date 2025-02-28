"use client";

import { useEffect, useState } from "react";
import Image from "next/image";
import { useParams } from "next/navigation";
import io from "socket.io-client";
import ExampleImg from "../../../public/exampleimg/room_img.jpg";

export default function RoomID() {
  const { id: roomId } = useParams();

  const hours = [
    "7:30 AM", "8:00 AM", "8:30 AM", "9:00 AM", "9:30 AM", "10:00 AM", "10:30 AM",
    "11:00 AM", "11:30 AM", "12:00 PM", "12:30 PM", "1:00 PM", "1:30 PM", "2:00 PM",
    "2:30 PM", "3:00 PM", "3:30 PM", "4:00 PM", "4:30 PM", "5:00 PM", "5:30 PM",
    "6:00 PM", "6:30 PM", "7:00 PM", "7:30 PM", "8:00 PM", "8:30 PM", "9:00 PM",
    "9:30 PM", "10:00 PM",
  ];

  const arr = [
    ["", "7:30 AM", "8:00 AM", "8:30 AM", "9:00 AM", "9:30 AM", "10:00 AM", "10:30 AM", "11:00 AM", "11:30 AM", "12:00 PM", "12:30 PM", "1:00 PM", "1:30 PM", "2:00 PM", "2:30 PM", "3:00 PM", "3:30 PM", "4:00 PM", "4:30 PM", "5:00 PM", "5:30 PM", "6:00 PM", "6:30 PM", "7:00 PM", "7:30 PM", "8:00 PM", "8:30 PM", "9:00 PM", "9:30 PM", "10:00 PM"],
    ["Monday", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
    ["Tuesday", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
    ["Wednesday", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
    ["Thursday", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
    ["Friday", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
    ["Saturday", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
    ["Sunday", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""],
  ];

  const [currentWeekOffset, setCurrentWeekOffset] = useState(0); // Offset tuần (0 là tuần hiện tại)
  const [orderDay, setOrderDay] = useState(getTodayDate());
  const [checkin, setCheckIn] = useState("");
  const [checkout, setCheckOut] = useState("");
  const [messages, setMessages] = useState<string[]>([]);
  const [socket, setSocket] = useState<SocketIOClient.Socket | null>(null);

  // Lấy ngày hiện tại định dạng YYYY-MM-DD
  function getTodayDate() {
    const today = new Date();
    return today.toISOString().split("T")[0];
  }

  // Giới hạn ngày trong tuần hiện tại
  const getWeekRange = () => {
    const today = new Date();
    const startOfWeek = new Date(today);
    startOfWeek.setDate(today.getDate() - today.getDay() + currentWeekOffset * 7);
    const endOfWeek = new Date(startOfWeek);
    endOfWeek.setDate(startOfWeek.getDate() + 6);
    return {
      min: startOfWeek.toISOString().split("T")[0],
      max: endOfWeek.toISOString().split("T")[0],
    };
  };

  // Chuyển tuần
  const prevWeek = () => setCurrentWeekOffset((prev) => prev - 1);
  const nextWeek = () => setCurrentWeekOffset((prev) => prev + 1);

  // Kiểm tra tuần hiện tại
  const isCurrentWeek = currentWeekOffset === 0;

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
    <div className="container flex flex-col min-h-screen mx-auto gap-8">
      <div className="flex flex-col w-full md:grid md:grid-cols-[55%_45%] gap-2">
        <div className="content-center font-extrabold text-2xl h-12 md:h-20 md:text-5xl md:col-span-2">
          ROOM NAME (ID: {roomId})
        </div>
        <div className="md:order-none md:h-[550px] h-96 content-center">
          <Image src={ExampleImg} alt="logo" className="h-auto max-h-full rounded-lg" />
        </div>
        <div className="flex flex-col gap-1 md:order-none md:h-auto min-h-96">
          <form onSubmit={sendRequest} className="flex flex-row items-center px-1 gap-2 h-12 border border-black">
            <label>Order day:</label>
            <input
              type="date"
              value={orderDay}
              onChange={(e) => setOrderDay(e.target.value)}
              min={getWeekRange().min}
              max={getWeekRange().max}
              className="h-8 border border-black"
            />
            <label>Check in:</label>
            <select
              value={checkin}
              onChange={(e) => setCheckIn(e.target.value)}
              className="h-8 border border-black"
            >
              <option value="">Select</option>
              {hours.map((hour) => (
                <option key={hour} value={hour}>{hour}</option>
              ))}
            </select>
            <label>Check out:</label>
            <select
              value={checkout}
              onChange={(e) => setCheckOut(e.target.value)}
              className="h-8 border border-black"
              disabled={!checkin}
            >
              <option value="">Select</option>
              {hours
                .filter((hour) => !checkin || hours.indexOf(hour) > hours.indexOf(checkin))
                .map((hour) => (
                  <option key={hour} value={hour}>{hour}</option>
                ))}
            </select>
            <button
              type="submit"
              className="p-2 w-24 bg-blue-500 text-white rounded border border-black"
            >
              Submit
            </button>
          </form>
          <div className="flex gap-2 my-2">
            <button
              onClick={prevWeek}
              disabled={isCurrentWeek}
              className="p-2 bg-gray-500 text-white rounded disabled:bg-gray-300"
            >
              Previous Week
            </button>
            <button
              onClick={nextWeek}
              className="p-2 bg-blue-500 text-white rounded"
            >
              Next Week
            </button>
          </div>
          <div className="flex-1">
            <div className="grid grid-cols-8 h-auto w-full border border-black">
              {arr.map((row, rowIndex) => (
                <div key={rowIndex} className="grid grid-rows-33 h-auto w-full">
                  {row.map((item, colIndex) => (
                    <div
                      key={`${rowIndex}-${colIndex}`}
                      className={`text-sm break-words text-center font-bold border border-black ${
                        colIndex > 0 && rowIndex > 0 && item === "1" ? "bg-red-300" :
                        colIndex > 0 && rowIndex > 0 && item === "2" ? "bg-green-300" : ""
                      }`}
                    >
                      {(colIndex === 0 || rowIndex === 0) && item}
                    </div>
                  ))}
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
      <div>
        <h2 className="text-xl font-bold mb-2">Room Request</h2>
        <ul className="border p-4 h-32 overflow-y-auto">
          {messages.map((msg, index) => (
            <li key={index} className="mb-2">{msg}</li>
          ))}
        </ul>
      </div>
    </div>
  );
}