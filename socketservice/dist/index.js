"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const app = (0, express_1.default)();
const port = 3001;
const http = require("http");
const server = http.createServer(app);
const { Server } = require("socket.io");
const io = new Server(server, {
    cors: {
        origin: "*",
    },
});
io.on("connect", (socket) => {
    console.log("Connected");
    socket.on("disconnect", () => {
        console.log("Disconnected");
    });
});
server.listen(port, () => {
    console.log(`⚡️[server]: Server is running at localhost:${port}`);
});
