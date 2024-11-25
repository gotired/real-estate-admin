import { useTheme } from "@mui/material";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import { ProgressBar } from "@/components/bar/ProgressBar";
import propertyReferralsInfo from "@/mock/propertyReferrals";
import DashboardCard from "../card/dashboard";

const PropertyReferrals = () => {
  const theme = useTheme();
  return (
    <DashboardCard flexDirection="column" gap="20px" height="100%">
      <Typography variant="h6" color={theme.palette.text.primary}>
        Property Referrals
      </Typography>
      <Stack direction="column" justifyContent="space-around" height="100%">
        {propertyReferralsInfo.map((bar) => (
          <ProgressBar
            key={bar.title}
            {...bar}
            bgBarColor="#555"
            fontColor={theme.palette.text.primary}
          />
        ))}
      </Stack>
    </DashboardCard>
  );
};

export default PropertyReferrals;
