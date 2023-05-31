import express, { Express, Request, Response } from "express";
import { Server as HTTPServer } from "http";
import { Socket, Server as SocketServer } from "socket.io";

const app: Express = express();
const port = 3001;
const http = require("http");
const server: HTTPServer = http.createServer(app);
const { Server } = require("socket.io");
const io: SocketServer = new Server(server);

const checkUserHasAccessToBoard = (
  Authorization_token: string,
  space_id: string
) => {
  return true;
};

io.on("connect", async (socket: Socket) => {
  let room_name = `${socket.handshake.query["space_id"]}-${socket.handshake.query["board_id"]}`;
  socket.join(room_name);
  socket.emit(
    "message",
    `Welcome ${socket.handshake.query["user"]}, to room [${room_name}] - From socketservice`
  );

  socket.on("remove_user", () => {});

  socket.on("save_board", () => {});

  socket.on("disconnect", () => {
    console.log(
      `🛑 socket disconnected - ${socket.handshake.query["space_id"]}-${socket.handshake.query["board_id"]}`
    );
  });

  console.log("connected", socket.handshake.query, socket.handshake.auth);
});

server.listen(port, () => {
  console.log(`⚡️[server]: Server is running at localhost:${port}`);
});
