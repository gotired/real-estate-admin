import { Box, Divider, Stack, Typography, useTheme } from "@mui/material";
import DashboardCard from "../card/dashboard";
import { BarChart } from "../charts/BarChart";

export const CustomerCard = () => {
  const theme = useTheme();
  const series = [
    {
      data: [183],
    },
    {
      data: [124],
    },
    {
      data: [507],
    },
    {
      data: [350],
    },
    {
      data: [199],
    },
  ];
  const colors = ["#CFC8FF", "#CFC8FF", "#475BE8", "#CFC8FF", "#CFC8FF"];
  return (
    <DashboardCard gap="20px" flexDirection="column" height="100%">
      <Typography variant="h6" color={theme.palette.text.primary}>
        Customer
      </Typography>
      <Stack
        display="flex"
        justifyContent="space-evenly"
        height="100%"
        gap="10px"
      >
        <Box display="flex" alignItems="center" justifyContent="space-between">
          <Stack height="fit-content" gap="12px">
            <Typography
              variant="subtitle2"
              color={theme.palette.text.secondary}
              noWrap
            >
              Total Customers
            </Typography>
            <Typography variant="h4">507K</Typography>
            <Typography color={theme.palette.primary.main}>21%</Typography>
          </Stack>
          <Box maxWidth="60%">
            <BarChart
              series={series}
              colors={colors}
              legend={[""]}
              ylabel={""}
              tooltipDetail={{
                formatter(val: number) {
                  return `${val}k$`;
                },
              }}
              options={{
                yaxis: {
                  show: false,
                },
                xaxis: {
                  show: false,
                },
                legend: { show: false },
                tooltip: {
                  enabled: false,
                },
              }}
              height="100%"
              width="100%"
            />
          </Box>
        </Box>
        <Divider />
        <Box display="flex" alignItems="center" justifyContent="space-between">
          <Stack height="fit-content" gap="12px">
            <Typography
              variant="subtitle2"
              color={theme.palette.text.secondary}
              noWrap
            >
              New Customers This Month
            </Typography>
            <Typography variant="h4">12K</Typography>
            <Typography color={theme.palette.primary.main}>21%</Typography>
          </Stack>
          <Box maxWidth="60%">
            <BarChart
              series={series}
              colors={colors}
              legend={[""]}
              ylabel={""}
              tooltipDetail={{
                formatter(val: number) {
                  return `${val}k$`;
                },
              }}
              options={{
                yaxis: {
                  show: false,
                },
                xaxis: {
                  show: false,
                },
                legend: { show: false },
                tooltip: {
                  enabled: false,
                },
              }}
              height="100%"
              width="85%"
            />
          </Box>
        </Box>
      </Stack>
    </DashboardCard>
  );
};
