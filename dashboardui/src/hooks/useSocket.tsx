import { useEffect, useMemo, useState } from "react";
import { Socket, io } from "socket.io-client";
import { getAuthToken } from "../helpers/authHelper";

export function useSocket(
  board_id: string | string[] | undefined,
  space_id: string | string[] | undefined,
  user: string | string[] | undefined
) {
  const [pageData, setPageData] = useState<{
    welcomeMessage: string;
    users: { [key: string]: { pos: { x: number; y: number }; color: string } };
  } | null>();
  const [users, setUsers] = useState<string[]>([]);
  const [userPositions, setUserPositions] = useState<{
    [key: string]: { pos: { x: number; y: number }; color: string };
  }>();
  const [message, setMessage] = useState<string>("");

  useEffect(() => {
    if (pageData && Object.keys(pageData?.users as any) != users) {
      setUsers(Object.keys(pageData?.users as any));
    }
    setUserPositions({ ...userPositions, ...pageData?.users });

    if (pageData?.welcomeMessage !== "") {
      setMessage(pageData?.welcomeMessage as string);
    }
  }, [pageData]);

  const [socket, setSocket] = useState<Socket>();

  useEffect(() => {
    socket &&
      socket.on("message", (recv: any) => {
        if (recv) {
          setPageData(JSON.parse(recv));
        }
      });

    socket?.on("board_positions", (recv: any) => {
      setPageData({ welcomeMessage: "", users: JSON.parse(recv) });
    });
  }, [socket]);

  useEffect(() => {
    if (board_id && space_id) {
      initialiseConnection(space_id as string, board_id as string);
    }

    return () => {
      disconnect();
    };
  }, [board_id, space_id]);

  function initialiseConnection(space_id: string, board_id: string) {
    let socket_url = process.env.NEXT_PUBLIC_SOCKET_API_URL;
    if (socket_url) {
      setSocket(
        io(socket_url, {
          reconnectionDelay: 10000,
          auth: {
            token: getAuthToken(),
          },
          query: {
            space_id,
            board_id,
            user,
          },
        })
      );
    }
  }

  function disconnect() {
    socket?.disconnect();
  }

  function UpdateUserPos(user: string, pos: any) {
    socket?.emit("updateuserpos", JSON.stringify({ user, pos }));
  }

  const memoedValue = useMemo(
    () => ({
      pageData,
      message,
      users,
      userPositions,
      UpdateUserPos,
    }),
    [pageData]
  );
  return memoedValue;
}
