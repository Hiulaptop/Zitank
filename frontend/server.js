"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var http_1 = require("http");
var url_1 = require("url");
var next_1 = require("next");
var ws_1 = require("ws");
var port = process.env.PORT ? parseInt(process.env.PORT) : 3000;
var dev = process.env.NODE_ENV !== "production";
var app = (0, next_1.default)({ dev: dev });
var handle = app.getRequestHandler();
app.prepare().then(function () {
    var server = (0, http_1.createServer)(function (req, res) {
        var parsedUrl = (0, url_1.parse)(req.url, true);
        handle(req, res, parsedUrl);
    });
    var wss = new ws_1.WebSocketServer({ server: server });
    wss.on("connection", function (ws) {
        console.log("Client connected!");
        ws.on("message", function (message) {
            try {
                var data = JSON.parse(message.toString());
                console.log("Received data:", data);
                var response = "Received: Order day: ".concat(data.orderDay, ", Check in: ").concat(data.checkin, ", Check out: ").concat(data.checkout);
                ws.send(response);
            }
            catch (error) {
                console.error("Error parsing data:", error);
                ws.send("Error: Invalid data");
            }
        });
        ws.on("close", function () {
            console.log("Client disconnected!");
        });
    });
    server.listen(port, function () {
        console.log("Server running on http://localhost:".concat(port, " and ws://localhost:").concat(port));
    });
});
