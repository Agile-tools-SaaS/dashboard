import Head from "next/head";

/**
 * homepage show all spaces calls get all spaces,
 * restricted if not logged in, this should show recent items,
 * (this will require a few things 5 most recent being stored
 * on the user account per space, an endpoint to return this
 * as well)
 */
export default function Home() {
  return (
    <>
      <Head>
        <title></title>
      </Head>
      <main>
        <p>Welcome back {"{user}"}</p>
      </main>
    </>
  );
}
