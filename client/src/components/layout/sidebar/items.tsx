import React from "react";
import { RxDashboard } from "react-icons/rx";
import { TbBuildingEstate } from "react-icons/tb";
import { SlPeople } from "react-icons/sl";
import { CiStar } from "react-icons/ci";
import { BiMessageDetail } from "react-icons/bi";
import { CgProfile } from "react-icons/cg";

export const pages = [
  {
    title: "Dashboard",
    icon: <RxDashboard />,
    path: "/",
  },
  { title: "Property", icon: <TbBuildingEstate />, path: "/property" },
  { title: "Agent", icon: <SlPeople />, path: "/agent" },
  { title: "Review", icon: <CiStar />, path: "/review" },
  { title: "Message", icon: <BiMessageDetail />, path: "/message" },
  { title: "My Profile", icon: <CgProfile />, path: "/profile" },
];
