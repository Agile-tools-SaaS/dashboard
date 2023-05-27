import { useEffect, useMemo, useState } from "react";
import { Socket, io } from "socket.io-client";
import { getAuthToken } from "../helpers/authHelper";

export function useSocket() {
  const [pageData, setPageData] = useState("");
  const [socket, setSocket] = useState<Socket>();

  useEffect(() => {
    socket &&
      socket.on("message", (recv: any) => {
        if (recv) {
          setPageData(recv);
        }
      });
  }, [socket]);

  function initialiseConnection(space_id: string, board_id: string) {
    let user = "callummcluskey100@gmail.com";
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

  const memoedValue = useMemo(
    () => ({
      initialiseConnection,
      pageData,
      disconnect,
    }),
    [pageData]
  );
  return memoedValue;
}
