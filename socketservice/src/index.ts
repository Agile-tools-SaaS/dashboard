import express, { Express, Request, Response } from "express";
import { Socket } from "socket.io";

const app: Express = express();
const port = 3001;
const http = require("http");
const server = http.createServer(app);
const { Server } = require("socket.io");
const io = new Server(server, {
  cors: {
    origin: "*",
  },
});

io.on("connect", (socket: any) => {
  console.log("Connected");
  socket.on("disconnect", () => {
    console.log("Disconnected");
  });
});

server.listen(port, () => {
  console.log(`⚡️[server]: Server is running at localhost:${port}`);
});
