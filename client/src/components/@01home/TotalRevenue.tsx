import { Box, Stack } from "@mui/material";
import Typography from "@mui/material/Typography";
import ArrowCircleUpRounded from "@mui/icons-material/ArrowCircleUpRounded";

import { useTheme } from "@mui/material";
import { BarChart } from "@/components/charts/BarChart";
import DashboardCard from "../card/dashboard";

export const TotalRevenue = () => {
  const theme = useTheme();

  const series = [
    {
      name: "Last Month",
      data: [183, 124, 115, 85, 143, 143, 96],
    },
    {
      name: "Running Month",
      data: [95, 84, 72, 44, 108, 108, 47],
    },
  ];
  const colors = ["#475BE8", "#CFC8FF"];
  const legend = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul"];

  return (
    <DashboardCard
      flexDirection="column"
      justifyContent="space-between"
      height="100%"
    >
      <Typography variant="h6" color={theme.palette.text.primary}>
        Total Revenue
      </Typography>
      <Stack direction="row" gap="25px" flexWrap="wrap" alignItems="center">
        <Typography
          variant="h5"
          color={theme.palette.text.primary}
          sx={{ fontWeight: "bold" }}
        >
          $236,535
        </Typography>
        <Stack direction="row" alignItems="center" gap={1}>
          <Box
            display="flex"
            justifyContent="center"
            alignItems="center"
            borderRadius="20px"
            height="20px"
            width="20px"
            sx={{ backgroundColor: theme.palette.primary.main }}
          >
            <ArrowCircleUpRounded
              sx={{
                fontSize: 25,
                // color: theme.palette.primary.main,
              }}
            />
          </Box>
          <Stack>
            <Typography fontSize={15} color={theme.palette.primary.main}>
              0.8%
            </Typography>
            <Typography variant="caption" color={theme.palette.text.secondary}>
              Than Last Month
            </Typography>
          </Stack>
        </Stack>
      </Stack>

      <BarChart
        series={series}
        colors={colors}
        legend={legend}
        ylabel={"$ (thousands)"}
        tooltipDetail={{
          formatter(val: number) {
            return `${val}k$`;
          },
        }}
      />
    </DashboardCard>
  );
};
