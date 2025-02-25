"use client"

import { use, useEffect, useState } from "react";
import { resolveMetadataItems } from "next/dist/lib/metadata/resolve-metadata";
import ExampleImg from "../../../public/exampleimg/room_img.jpg"
import Image from "next/image";


export default function RoomID() {

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

    const hours = [
        "7:30 AM", "8:00 AM", "8:30 AM", "9:00 AM", "9:30 AM",
        "10:00 AM", "10:30 AM", "11:00 AM", "11:30 AM", "12:00 PM",
        "12:30 PM", "1:00 PM", "1:30 PM", "2:00 PM", "2:30 PM",
        "3:00 PM", "3:30 PM", "4:00 PM", "4:30 PM", "5:00 PM",
        "5:30 PM", "6:00 PM", "6:30 PM", "7:00 PM", "7:30 PM",
        "8:00 PM", "8:30 PM", "9:00 PM", "9:30 PM", "10:00 PM"
    ];

    const [orderDay, setOrderDay] = useState("1970-01-01"); // Giá trị mặc định hợp lệ cho input date
    const [checkin, setCheckIn] = useState("");
    const [checkout, setCheckOut] = useState("");
    const [received, setReceived] = useState(""); // Thêm state để hiển thị dữ liệu nhận được
    const [ws, setWs] = useState<WebSocket | null>(null);

    useEffect(() => {
        const socket = new WebSocket("ws://localhost:3000");

        socket.onopen = () => {
            console.log("WebSocket connected.");
        };

        socket.onmessage = (event) => {
            console.log("Received:", event.data);
            setReceived(event.data); // Cập nhật state khi nhận dữ liệu
        };

        socket.onerror = (error) => {
            console.error("WebSocket error:", error);
        };

        socket.onclose = () => {
            console.log("WebSocket closed!");
        };

        setWs(socket);

        return () => {
            socket.close();
        };
    }, []);

    const sendOrder = (e: React.FormEvent) => {
        e.preventDefault();
        if (ws && ws.readyState === WebSocket.OPEN && orderDay && checkin && checkout) {
            const data = { orderDay, checkin, checkout };
            ws.send(JSON.stringify(data));
        } else {
            console.error("WebSocket is not open or data is missing.");
        }
    };

    return (
        <div className="container flex flex-col min-h-screen mx-auto gap-8">
            {/* Room order */}
            <div className="flex flex-col w-full md:grid md:grid-cols-[55%_45%] gap-2">
                <div className="content-center font-extrabold text-2xl h-12 md:h-20 md:text-5xl md:col-span-2">
                    ROOM NAME
                </div>
                <div className="md:order-none md:h-[550px] h-96 content-center">
                    <Image src={ExampleImg} alt="logo" className="h-auto max-h-full rounded-lg" />
                </div>
                <div className="flex flex-col gap-1 md:order-none md:h-auto min-h-96">
                    {/* ORDER REQUEST */}
                    <form className="flex flex-row items-center px-1 gap-1 h-12 content-center border border-black">
                        <label >Order day:</label>
                        <input type="date" id="order-day" name="order-day" className="h-8 border border-black" value={orderDay} onChange={(e) => setOrderDay(e.target.value)} />

                        <label >Check in:</label>
                        <select name="check-in" id="check-in" className="h-8 border border-black" value={checkin} onChange={(e) => setCheckIn(e.target.value)}>
                            {hours.map((hours, index) => (
                                <option value={`check-in-${index}`}>{hours}</option>
                            ))}
                        </select>

                        <label >Check out:</label>
                        <select name="check-out" id="check-out" className="h-8 border border-black" value={checkout} onChange={(e) => setCheckOut(e.target.value)}>
                            {hours.map((hours, index) => (
                                <option value={`check-out-${index}`}>{hours}</option>
                            ))}
                        </select>

                        <button onClick={sendOrder} type="submit" className="mr-0 justify-self-end h-8 w-24 border border-black">
                            Order
                        </button>
                    </form>

                    <p>Received: {received}</p>


                    <div className="flex-1 ">
                        <div className="grid grid-cols-8 h-full w-full border border-black">
                            {arr.map((row, rowIndex) => (
                                <div key={rowIndex} className="grid grid-rows-[33] h-full w-full ">
                                    {row.map((item, colIndex) => (
                                        <div key={`${rowIndex}-${colIndex}`}
                                            className={`
                                                ${colIndex > 0 && rowIndex > 0 && item === '1' && 'bg-red-300'}
                                                ${colIndex > 0 && rowIndex > 0 && item === '2' && 'bg-green-300'}
                                                text-sm break-words text-center font-bold border border-black
                                            `}
                                        >
                                            {(colIndex == 0 || rowIndex == 0) && item}
                                            {/* {`${rowIndex * 31 + colIndex}`} */}
                                        </div>
                                    ))}
                                </div>
                            ))}
                        </div>

                    </div>


                </div>
            </div>

            {/* Other rooms */}
            <div>

            </div>
        </div>
    );
}
