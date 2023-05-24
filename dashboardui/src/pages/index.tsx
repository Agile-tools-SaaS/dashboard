import Head from "next/head";
import Link from "next/link";

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
        <title>Floboard</title>
      </Head>
      <main>
        <h1>Floboard</h1>
        <p>Welcome back {"{user}"}</p>
        <hr />
        <p>site map</p>
        <div style={{ display: "flex", flexDirection: "column", gap: 4 }}>
          <Link href="/login">login page</Link>
          <Link href="/user/account">account page</Link>
          <Link href="/space/example_space">example_space space</Link>
          <Link href="/space/example_space/board/example_board">
            example_board board inside of example_space space
          </Link>
          <Link href="/404">error page</Link>
        </div>
      </main>
    </>
  );
}
