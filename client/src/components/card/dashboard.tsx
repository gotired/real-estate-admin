import { Box, BoxProps, useTheme } from "@mui/material";
import { ReactNode } from "react";

interface DashboardCardProps extends BoxProps {
  children: ReactNode;
}

export default function DashboardCard({
  children,
  ...props
}: DashboardCardProps) {
  const theme = useTheme();

  return (
    <Box
      bgcolor={theme.palette.background.default}
      display="flex"
      padding="25px"
      borderRadius="15px"
      {...props}
    >
      {children}
    </Box>
  );
}
