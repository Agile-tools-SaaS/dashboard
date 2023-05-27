import express, { Express, Request, Response } from "express";
import { Server as HTTPServer } from "http";
import { Socket, Server as SocketServer } from "socket.io";

const app: Express = express();
const port = 3001;
const http = require("http");
const server: HTTPServer = http.createServer(app);
const { Server } = require("socket.io");
const io: SocketServer = new Server(server, {
  cors: {
    origin: "*",
  },
});

const getBoard = (board_id: string, space_id: string): Promise<any> => {
  return new Promise(() => true);
};

const onBoardConnect = (socket: Socket) => {
  console.log("connected");
  socket.on("disconnect", () => {});
  socket.on("remove_user", () => {});
  socket.on("save_board", () => {});
};

const checkUserHasAccessToBoard = (
  Authorization_token: string
): Promise<any> => {
  return new Promise(() => true);
};

io.on("connect", async (socket: Socket) => {
  // endpoint to check if user has access to the server
  // let has_access = await checkUserHasAccessToBoard(
  //   socket.request.headers["Authorization"] as string
  // );

  socket.join(
    `${socket.handshake.query["space_id"]}-${socket.handshake.query["board_id"]}`
  );

  socket.emit("message", "Welcome");

  socket.on("disconnect", () => {
    console.log(
      `🛑 socket disconnected - ${socket.handshake.query["space_id"]}-${socket.handshake.query["board_id"]}`
    );
  });

  console.log("connected", socket.handshake.query, socket.handshake.auth);

  let has_access = true;

  // if (has_access) {
  //   // join user
  //   // check if users are in board
  //   let boardUsers = [];

  //   if (boardUsers.length === 0) {
  //     getBoard("", "").then(() => {
  //       onBoardConnect(socket);
  //     });
  //   } else {
  //     onBoardConnect(socket);
  //   }
  // } else {
  //   // reject the request
  // }
});

server.listen(port, () => {
  console.log(`⚡️[server]: Server is running at localhost:${port}`);
});
