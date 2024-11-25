"use client";

import { TotalRevenue } from "@/components/@01home/TotalRevenue";
import PieChart from "@/components/@01home/PieChartCard";
import PropertyReferrals from "@/components/@01home/PropertyReferrals";
import pieChartMock from "@/mock/pieChart";
import { Typography } from "@mui/material";
import Grid from "@mui/material/Grid2";
import PageCard from "@/components/card/page";
import { TopAgent } from "@/components/@01home/TopAgent";
import { CustomerCard } from "@/components/@01home/CustomerCard";
import { LatestSalesCard } from "@/components/@01home/LatestSalesCard";

export default function Home() {
  return (
    <PageCard>
      <Typography variant="h5">Dashboard</Typography>
      <Grid container spacing="25px">
        {pieChartMock.map((value, index) => (
          <Grid size={3} key={index}>
            <PieChart
              title={value.title}
              value={value.value}
              series={value.series}
              colors={value.colors}
            />
          </Grid>
        ))}
      </Grid>
      <Grid container spacing="25px">
        <Grid size={8}>
          <TotalRevenue />
        </Grid>
        <Grid size={4}>
          <PropertyReferrals />
        </Grid>
      </Grid>
      <Grid container spacing="25px">
        <Grid size={4}>
          <TopAgent />
        </Grid>
        <Grid size={4}>
          <CustomerCard />
        </Grid>
        <Grid size={4}>
          <LatestSalesCard />
        </Grid>
      </Grid>
    </PageCard>
  );
}
