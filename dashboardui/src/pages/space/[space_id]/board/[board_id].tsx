import { Pointer } from "@/components/mouse/pointer";
import { useMousePos } from "@/hooks/useMousePos";
import { useSocket } from "@/hooks/useSocket";
import Head from "next/head";
import { useRouter } from "next/router";
import "@/styles/board/main.module.scss";
import { BoardContainer } from "@/components/board/container";
import { useEffect, useRef, useState } from "react";
import { useInterval } from "@/hooks/useInterval";
import { motion } from "framer-motion";
import { UsersBox } from "@/components/board/userBox";

interface mouseCoordinates {
  x: number;
  y: number;
}
export default function BoardView() {
  const router = useRouter();
  let { board_id, space_id, user } = router.query;

  const { users, userPositions, pageData, UpdateUserPos, message } = useSocket(
    board_id,
    space_id,
    user
  );

  // let mousePos = useMousePos();

  const [mousePos, setMousePos] = useState<mouseCoordinates>({ x: 0, y: 0 });

  useEffect(() => {
    const handleMouseMove = (e: MouseEvent) => {
      setMousePos({ x: e.clientX, y: e.clientY });
    };

    window.addEventListener("mousemove", handleMouseMove);

    return () => {
      window.removeEventListener("mousemove", handleMouseMove);
    };
  }, []);

  useInterval(() => {
    UpdateUserPos(user as string, mousePos);
  }, 300);

  return (
    <>
      <Head>
        <title></title>
      </Head>
      <main>
        <p>{message}</p>

        {userPositions &&
          users.map(
            (x, i) =>
              x !== user && (
                <Pointer
                  otherUser={true}
                  coordinates={userPositions[x].pos ?? { x: 50, y: 50 }}
                  user={users[i]}
                  color={userPositions[x].color}
                />
              )
          )}

        {userPositions && (
          <UsersBox
            {...users.map((x) => ({
              name: x as string,
              color: userPositions[x].color as string,
            }))}
          />
        )}

        <BoardContainer>
          <>
            {userPositions && (
              <Pointer
                coordinates={mousePos}
                color={userPositions[user as string]?.color}
                user={user as string}
              />
            )}
          </>
        </BoardContainer>
      </main>
    </>
  );
}
