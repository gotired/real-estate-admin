import { ReactNode } from "react";
import "../styles/globals.css";
import { Stack } from "@mui/material";
import { Header } from "./layout/header";
import { SideBar } from "./layout/sidebar";

export default function Layout({ children }: { children: ReactNode }) {
  return (
    <main>
      <Stack maxHeight="100vh">
        <Header />
        <Stack flex={1} direction="row" maxHeight="calc(100vh - 64px)">
          <SideBar />
          {children}
        </Stack>
      </Stack>
    </main>
  );
}
