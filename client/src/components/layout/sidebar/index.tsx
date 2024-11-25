import React from "react";
import Stack from "@mui/material/Stack";
import { useTheme } from "@mui/material";
import { RxDashboard } from "react-icons/rx";
import { TbBuildingEstate } from "react-icons/tb";
import { SlPeople } from "react-icons/sl";
import { CiStar } from "react-icons/ci";
import { BiMessageDetail } from "react-icons/bi";
import { CgProfile } from "react-icons/cg";
import { SideBarIcon } from "./sidebarItem";
import { pages } from "./items";

export const SideBar: React.FC = () => {
  const theme = useTheme();

  return (
    <Stack
      display="flex"
      padding="25px 16px"
      width="250px"
      bgcolor={theme.palette.background.default}
      maxHeight="500px"
    >
      {pages.map((page, index) => (
        <SideBarIcon
          title={page.title}
          icon={page.icon}
          path={page.path}
          key={index}
        />
      ))}
    </Stack>
  );
};
