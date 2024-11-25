import { Avatar, Box, Stack, Typography, useTheme } from "@mui/material";
import DashboardCard from "../card/dashboard";
import { FiMoreVertical } from "react-icons/fi";

interface AgentDetailProp {
  name: string;
  title: string;
}

const AgentDetail = ({ name, title }: AgentDetailProp) => {
  const theme = useTheme();
  return (
    <Box display="flex" flexDirection="row" gap="10px" alignItems="center">
      <Avatar>LI</Avatar>
      <Stack width="100%">
        <Stack
          display="flex"
          flexDirection="row"
          justifyContent="space-between"
          alignItems="center"
        >
          <Typography variant="subtitle1">{name}</Typography>
          <FiMoreVertical height="100%" />
        </Stack>
        <Typography variant="subtitle2" color={theme.palette.text.secondary}>
          {title}
        </Typography>
      </Stack>
    </Box>
  );
};

export const TopAgent = () => {
  const theme = useTheme();
  return (
    <DashboardCard gap="20px" flexDirection="column">
      <Typography variant="h6" color={theme.palette.text.primary}>
        Top Agent
      </Typography>
      <Stack flex={1} justifyContent="space-evenly" gap="20px">
        <AgentDetail name="Lorem Ipsum" title="Top Agent" />
        <AgentDetail name="Lorem Ipsum" title="Top Agent" />
        <AgentDetail name="Lorem Ipsum" title="Top Agent" />
        <AgentDetail name="Lorem Ipsum" title="Top Agent" />
        <AgentDetail name="Lorem Ipsum" title="Top Agent" />
      </Stack>
    </DashboardCard>
  );
};
