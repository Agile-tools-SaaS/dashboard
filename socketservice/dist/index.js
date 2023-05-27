"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
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
const getBoard = (board_id, space_id) => {
    return new Promise(() => true);
};
const onBoardConnect = (socket) => {
    console.log("connected");
    socket.on("disconnect", () => { });
    socket.on("remove_user", () => { });
    socket.on("save_board", () => { });
};
const checkUserHasAccessToBoard = (Authorization_token) => {
    return new Promise(() => true);
};
io.on("connect", (socket) => __awaiter(void 0, void 0, void 0, function* () {
    // endpoint to check if user has access to the server
    // let has_access = await checkUserHasAccessToBoard(
    //   socket.request.headers["Authorization"] as string
    // );
    socket.join(`${socket.handshake.query["space_id"]}-${socket.handshake.query["board_id"]}`);
    socket.emit("message", "Welcome");
    socket.on("disconnect", () => {
        console.log(`🛑 socket disconnected - ${socket.handshake.query["space_id"]}-${socket.handshake.query["board_id"]}`);
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
}));
server.listen(port, () => {
    console.log(`⚡️[server]: Server is running at localhost:${port}`);
});
