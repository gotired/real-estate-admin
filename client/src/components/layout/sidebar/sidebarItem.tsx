import React from "react";
import Typography from "@mui/material/Typography";
import { Box, Link, useTheme } from "@mui/material";
import { usePathname } from "next/navigation";
import { sideBarItemProp } from "@/interfaces/sidebar";

export const SideBarIcon = ({ title, icon, path }: sideBarItemProp) => {
  const theme = useTheme();
  const currentPath = usePathname();

  return (
    <Link href={path} underline="none">
      <Box
        paddingY="16px"
        paddingX="24px"
        sx={{
          bgcolor:
            currentPath === path ? theme.palette.primary.main : undefined,
          borderRadius: "12px",
          display: "flex",
          flexDirection: "row",
          alignItems: "center",
          height: "56px",
          gap: "10px",
          color:
            currentPath == path
              ? theme.palette.text.primary
              : theme.palette.text.secondary,
          ":hover": {
            color: theme.palette.text.primary,
          },
        }}
      >
        <Box>{icon}</Box>
        <Typography variant="body1">{title}</Typography>
      </Box>
    </Link>
  );
};
