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

const colors = [
  "#003d5b",
  "#91972A",
  "#d1495b",
  "#F5853F",
  "#40531B",
  "#C81D25",
];

let users: any = {};

io.on("connect", async (socket: Socket) => {
  let room_name = `${socket.handshake.query["space_id"]}-${socket.handshake.query["board_id"]}`;
  let user_name: string = socket.handshake.query["user"] as string;
  socket.join(room_name);
  users[user_name] = {
    pos: { x: 0, y: 0 },
    color: colors[Math.floor(Math.random() * colors.length)],
  };
  socket.emit(
    "message",
    JSON.stringify({
      welcomeMessage: `Welcome ${socket.handshake.query["user"]}, to room [${room_name}] - From socketservice`,
      users,
    })
  );

  socket.on("updateuserpos", (recv) => {
    let updated_user = JSON.parse(recv);

    users[updated_user.user] = {
      pos: updated_user.pos,
      color: users[updated_user.user].color,
    };
    socket.emit("board_positions", JSON.stringify(users));
  });

  socket.on("change", () => {});

  socket.on("remove_user", () => {});

  socket.on("save_board", () => {});

  socket.on("disconnect", () => {
    delete users[user_name];
    console.log(
      `🛑 socket disconnected - ${socket.handshake.query["space_id"]}-${socket.handshake.query["board_id"]}`
    );
  });

  console.log("connected", socket.handshake.query, socket.handshake.auth);
});

server.listen(port, () => {
  console.log(`⚡️[server]: Server is running at localhost:${port}`);
});
