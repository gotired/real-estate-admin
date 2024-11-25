import { Stack, useTheme } from "@mui/material";
import { ReactNode } from "react";

export default function PageCard({ children }: { children: ReactNode }) {
  const theme = useTheme();
  return (
    <Stack
      bgcolor={theme.palette.background.paper}
      flexGrow="1"
      flex={1}
      direction="column"
      padding="25px"
      gap="25px"
      overflow="auto"
    >
      {children}
    </Stack>
  );
}
