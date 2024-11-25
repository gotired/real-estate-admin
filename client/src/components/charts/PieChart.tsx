import { ApexOptions } from "apexcharts";
import { ApexChart } from "./ApexChart";
import merge from "lodash/merge";

export const PieChart = ({
  series,
  colors,
  options = {},
  ...prop
}: PieChartProps) => {
  const defaultOptions: ApexOptions = {
    chart: { type: "donut" },
    colors,
    legend: { show: false },
    dataLabels: {
      enabled: false,
    },
    tooltip: {
      enabled: false,
    },
    stroke: {
      show: false,
    },
    plotOptions: {
      pie: {
        expandOnClick: false,
        donut: {
          size: "55%",
        },
      },
    },
  };

  const mergedOptions = merge({}, defaultOptions, options);
  return (
    <ApexChart options={mergedOptions} series={series} type="donut" {...prop} />
  );
};
