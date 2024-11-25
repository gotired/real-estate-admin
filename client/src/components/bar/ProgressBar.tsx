import { Box, Stack, Typography } from "@mui/material";
import React from "react";

interface ProgressBarProps {
  title: string;
  percentage: number;
  color: string;
  bgBarColor: string;
  fontColor: string;
}

export const ProgressBar: React.FC<ProgressBarProps> = ({
  title,
  percentage,
  color,
  bgBarColor,
  fontColor,
}) => (
  <Box width="100%">
    <Stack direction="row" alignItems="center" justifyContent="space-between">
      <Typography fontSize={16} fontWeight={500} color={fontColor}>
        {title}
      </Typography>
      <Typography fontSize={16} fontWeight={500} color={fontColor}>
        {percentage}%
      </Typography>
    </Stack>
    <Box
      mt={2}
      position="relative"
      width="100%"
      height="8px"
      borderRadius={1}
      bgcolor={bgBarColor}
    >
      <Box
        width={`${percentage}%`}
        bgcolor={color}
        position="absolute"
        height="100%"
        borderRadius={1}
      />
    </Box>
  </Box>
);
