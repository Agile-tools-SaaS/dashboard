import { useSocket } from "@/hooks/useSocket";
import Head from "next/head";
import { useRouter } from "next/router";
import { useEffect } from "react";

export default function BoardView() {
  const router = useRouter();

  let { board_id, space_id } = router.query;
  const { initialiseConnection, pageData, disconnect } = useSocket();

  useEffect(() => {
    if (board_id && space_id) {
      initialiseConnection(board_id as string, space_id as string);
    }

    return () => {
      disconnect();
    };
  }, [board_id, space_id]);

  return (
    <>
      <Head>
        <title></title>
      </Head>
      <main>
        <p> board: {board_id}</p>
        <p> space: {space_id}</p>
        <p>{JSON.stringify(pageData)}</p>
      </main>
    </>
  );
}
