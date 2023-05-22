import Head from "next/head";
import { useRouter } from "next/router";

export default function SpaceDashboard() {
  const router = useRouter();

  let { space_id } = router.query;

  return (
    <>
      <Head>
        <title></title>
      </Head>
      <main>
        <p> space: {space_id}</p>
      </main>
    </>
  );
}
