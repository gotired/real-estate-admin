import Typography from "@mui/material/Typography";
import Stack from "@mui/material/Stack";

import { useTheme } from "@mui/material";
import { PieChart } from "@/components/charts/PieChart";
import DashboardCard from "../card/dashboard";

const PieChartCard = ({ title, value, series, colors }: PieChartCardProps) => {
  const theme = useTheme();
  return (
    <DashboardCard
      flexDirection="row"
      justifyContent="space-between"
      alignItems="center"
      display="flex"
    >
      <Stack direction="column" width="fit-content">
        <Typography fontSize={14} color={theme.palette.text.secondary}>
          {title}
        </Typography>
        <Typography
          fontSize={24}
          color={theme.palette.text.primary}
          fontWeight={700}
          mt={1}
        >
          {value}
        </Typography>
      </Stack>
      <PieChart series={series} colors={colors} width="100px" />
    </DashboardCard>
  );
};

export default PieChartCard;
