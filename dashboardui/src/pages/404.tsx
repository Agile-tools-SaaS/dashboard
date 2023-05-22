import { Head } from "next/document";

export default function Custom404() {
  return (
    <>
      <Head>
        <title>Uh oh</title>
      </Head>
      <main>
        <h1>404</h1>
        <p>This page does not exist.</p>
      </main>
    </>
  );
}
