import React from "react";
import Typography from "@mui/material/Typography";
import { Box } from "@mui/material";
import Image from "next/image";
import logo from "@/assets/logo.svg";
import Link from "next/link";

export const Title: React.FC = () => {
  return (
    <Box width="250px">
      <Box display="flex" alignItems="center">
        <Link href="/">
          <Image src={logo} alt=""></Image>
        </Link>
        <Link href="/">
          <Typography variant="h5" paddingLeft="10px">
            Property
          </Typography>
        </Link>
      </Box>
    </Box>
  );
};
