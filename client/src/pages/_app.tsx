import type { AppProps } from "next/app";
import Layout from "@/components/layout";
import { ColorModeContextProvider } from "@/contexts";
import { CssBaseline, GlobalStyles } from "@mui/material";
import "../styles/globals.css";

export default function MyApp({ Component, pageProps }: AppProps) {
  return (
    <ColorModeContextProvider>
      <Layout>
        <CssBaseline />
        <GlobalStyles styles={{ html: { WebkitFontSmoothing: "auto" } }} />
        <Component {...pageProps} />
      </Layout>
    </ColorModeContextProvider>
  );
}
