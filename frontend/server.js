// // "use strict";
// // Object.defineProperty(exports, "__esModule", { value: true });
// // var http_1 = require("http");
// // var url_1 = require("url");
// // var next_1 = require("next");
// // var ws_1 = require("ws");
// // var port = process.env.PORT ? parseInt(process.env.PORT) : 3000;
// // var dev = process.env.NODE_ENV !== "production";
// // var app = (0, next_1.default)({ dev: dev });
// // var handle = app.getRequestHandler();
// // app.prepare().then(function () {
// //     var server = (0, http_1.createServer)(function (req, res) {
// //         var parsedUrl = (0, url_1.parse)(req.url, true);
// //         handle(req, res, parsedUrl);
// //     });
// //     var wss = new ws_1.WebSocketServer({ server: server });
// //     wss.on("connection", function (ws) {
// //         console.log("Client connected!");
// //         ws.on("message", function (message) {
// //             try {
// //                 var data = JSON.parse(message.toString());
// //                 console.log("Received data:", data);
// //                 var response = "Received: Order day: ".concat(data.orderDay, ", Check in: ").concat(data.checkin, ", Check out: ").concat(data.checkout);
// //                 ws.send(response);
// //             }
// //             catch (error) {
// //                 console.error("Error parsing data:", error);
// //                 ws.send("Error: Invalid data");
// //             }
// //         });
// //         ws.on("close", function () {
// //             console.log("Client disconnected!");
// //         });
// //     });
// //     server.listen(port, function () {
// //         console.log("Server running on http://localhost:".concat(port, " and ws://localhost:").concat(port));
// //     });
// // });

// const { createServer } = require("http");
// const { parse } = require("url");
// const next = require("next");
// const { Server } = require("socket.io");

// const dev = process.env.NODE_ENV !== "production";
// const port = process.env.PORT || 3000;
// const app = next({ dev });
// const handle = app.getRequestHandler();

// async function startServer() {
//   try {
//     await app.prepare();
//     const server = createServer((req, res) => {
//       const parsedUrl = parse(req.url, true);
//       handle(req, res, parsedUrl);
//     });

//     const io = new Server(server);

//     io.on("connection", (socket) => {
//       console.log("New client connected:", socket.id);

//       // Xử lý khi client tham gia room dựa trên test_id
//       socket.on("joinRoom", (roomId) => {
//         socket.join(roomId);
//         console.log(`Client ${socket.id} joined room: ${roomId}`);
//       });

//       // Xử lý tin nhắn trong room
//       socket.on("message", ({ roomId, message }) => {
//         console.log(`Message in room ${roomId}: ${message}`);
//         io.to(roomId).emit("message", message); // Gửi tin nhắn chỉ trong room
//       });

//       socket.on("disconnect", () => {
//         console.log("Client disconnected:", socket.id);
//       });
//     });

//     server.listen(port, () => {
//       console.log(`> Server running on http://localhost:${port}`);
//     });
//   } catch (err) {
//     console.error("Failed to start server:", err);
//     process.exit(1);
//   }
// }

// startServer();


const { createServer } = require("http");
const { parse } = require("url");
const next = require("next");
const { Server } = require("socket.io");

const dev = process.env.NODE_ENV !== "production";
const port = process.env.PORT || 3000;
const app = next({ dev });
const handle = app.getRequestHandler();

async function startServer() {
  try {
    await app.prepare();
    const server = createServer((req, res) => {
      const parsedUrl = parse(req.url, true);
      handle(req, res, parsedUrl);
    });

    const io = new Server(server, {
      cors: { origin: "*" }, // Cho phép kết nối từ localhost:3000
    });

    io.on("connection", (socket) => {
      console.log("Client connected:", socket.id);

      socket.on("joinRoom", (roomId) => {
        socket.join(roomId);
        console.log(`Client ${socket.id} joined room ${roomId}`);
      });

      socket.on("message", ({ roomId, message }) => {
        console.log(`Message in room ${roomId}: ${message}`);
        io.to(roomId).emit("message", message);
      });

      socket.on("disconnect", () => {
        console.log("Client disconnected:", socket.id);
      });
    });

    server.listen(port, () => {
      console.log(`> Server running on http://localhost:${port}`);
    });
  } catch (error) {
    console.error("Server error:", error);
    process.exit(1);
  }
}

startServer();