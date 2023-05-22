import Head from "next/head";
import { useRouter } from "next/router";

export default function BoardView() {
  const router = useRouter();

  let { board_id, space_id } = router.query;
  return (
    <>
      <Head>
        <title></title>
      </Head>
      <main>
        <p> board: {board_id}</p>
        <p> space: {space_id}</p>
      </main>
    </>
  );
}
