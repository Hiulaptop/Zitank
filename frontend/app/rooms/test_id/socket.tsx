import { WebSocketServer } from "ws";
import { NextApiRequest, NextApiResponse } from "next";

export default function handler(req: NextApiRequest, res: NextApiResponse) {
  const socket = res.socket as any; // Type assertion to avoid TypeScript error
  if (!socket.server.wss) {
    console.log("Initializing WebSocket server...");
    const wss = new WebSocketServer({ server: socket.server });

    wss.on("connection", (ws) => {
      console.log("Client connected!");
      ws.on("message", (message) => {
        try {
          const data = JSON.parse(message.toString());
          console.log("Received data:", data);
          const response = `Received: Order day: ${data.orderDay}, Check in: ${data.checkin}, Check out: ${data.checkout}`;
          ws.send(response);
        } catch (error) {
          console.error("Error parsing data:", error);
          ws.send("Error: Invalid data");
        }
      });
      ws.on("close", () => {
        console.log("Client disconnected!");
      });
    });

    socket.server.wss = wss; // Attach WebSocket server to the server
  }

  res.end();
}

export const config = {
  api: {
    bodyParser: false,
  },
};