import { MenuBar } from "@/components/general/menubar";
import { AuthProvider } from "@/hooks/useAuth";
import "@/styles/globals.css";
import type { AppProps } from "next/app";

export default function App({ Component, pageProps }: AppProps) {
  return (
    <AuthProvider>
      <MenuBar />
      <Component {...pageProps} />
    </AuthProvider>
  );
}
