import { Box, Stack, Typography, useTheme } from "@mui/material";
import DashboardCard from "../card/dashboard";
import Image, { StaticImageData } from "next/image";
import React from "react";
import salesCardMock from "@/mock/salesProperty";

interface SalesCardProp {
  image: StaticImageData | string;
  width: number;
  height: number;
  name: string;
  location: string;
  price: string;
}

const PropertyItem: React.FC<SalesCardProp> = ({
  image,
  width,
  height,
  name,
  location,
  price,
}) => {
  const theme = useTheme();

  return (
    <Stack
      display="flex"
      direction="row"
      justifyContent="space-between"
      alignItems="center"
    >
      <Stack display="flex" direction="row" alignItems="center" gap="10px">
        <Box
          sx={{
            width: width,
            height: height,
            overflow: "hidden",
            borderRadius: "5px",
            position: "relative",
          }}
        >
          <Image
            src={image}
            alt=""
            fill
            style={{
              objectFit: "cover",
            }}
          />
        </Box>
        <Stack gap="5px">
          <Typography variant="subtitle1">{name}</Typography>
          <Typography variant="subtitle2" color={theme.palette.text.secondary}>
            {location}
          </Typography>
        </Stack>
      </Stack>
      <Typography variant="h6">{price}</Typography>
    </Stack>
  );
};

export const LatestSalesCard = () => {
  const theme = useTheme();
  return (
    <DashboardCard gap="20px" flexDirection="column" height="100%">
      <Typography variant="h6" color={theme.palette.text.primary}>
        Latest Sales
      </Typography>
      <Stack
        display="flex"
        justifyContent="space-evenly"
        height="100%"
        gap="20px"
      >
        {salesCardMock.map((val, index) => (
          <PropertyItem
            image={val.image}
            width={60}
            height={60}
            name={val.name}
            location={val.location}
            price={val.price}
            key={index}
          />
        ))}
      </Stack>
    </DashboardCard>
  );
};
