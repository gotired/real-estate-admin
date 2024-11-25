import React from "react";
import AppBar from "@mui/material/AppBar";
import Stack from "@mui/material/Stack";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import { Avatar, Box, InputBase, useTheme } from "@mui/material";
import { Title } from "../title";

export const Header: React.FC = () => {
  const theme = useTheme();
  return (
    <AppBar
      color="default"
      position="sticky"
      elevation={0}
      sx={{ background: theme.palette.background.default }}
    >
      <Toolbar>
        <Stack
          direction="row"
          width="100%"
          justifyContent="space-between"
          alignItems="center"
        >
          <Stack flexDirection="row" justifyContent="center">
            <Title />
            <Box>
              <InputBase
                sx={{
                  height: "38px",
                  width: "405px",
                  border: 1,
                  borderRadius: "10px",
                  padding: "11px 12px",
                  color: theme.palette.text.primary,
                  borderColor: theme.palette.divider,
                  "&.Mui-focused": {
                    borderColor: theme.palette.text.secondary,
                  },
                  fontSize: "12px",
                }}
                type="email"
                placeholder="Search Property, Customer etc"
              />
            </Box>
          </Stack>
          <Stack
            display="flex"
            flexDirection="row"
            alignItems="center"
            gap="10px"
          >
            <Avatar>LI</Avatar>
            <Stack display="flex" gap="2px">
              <Typography variant="subtitle2">Lorem Ipsum</Typography>
            </Stack>
          </Stack>
        </Stack>
      </Toolbar>
    </AppBar>
  );
};
