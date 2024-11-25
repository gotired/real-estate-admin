import { ApexChart } from "./ApexChart";
import { ApexOptions } from "apexcharts";
import merge from "lodash/merge";

export const BarChart = ({
  series,
  colors,
  legend,
  ylabel,
  tooltipDetail,
  options = {},
  ...prop
}: BarChartComponentProp) => {
  const defaultOptions: ApexOptions = {
    chart: {
      type: "bar",
      toolbar: {
        show: false,
      },
    },
    colors,
    plotOptions: {
      bar: {
        borderRadius: 4,
        horizontal: false,
        columnWidth: "55%",
      },
    },
    dataLabels: {
      enabled: false,
    },
    grid: {
      show: false,
    },
    stroke: {
      colors: ["transparent"],
      width: 4,
    },
    xaxis: {
      categories: legend,
    },
    yaxis: {
      title: {
        text: ylabel,
      },
    },
    fill: {
      opacity: 1,
    },
    legend: {
      position: "top",
      horizontalAlign: "right",
    },
    tooltip: {
      theme: "true",
      custom: function ({ series, seriesIndex, dataPointIndex, w }) {
        return (
          '<div style="background: #555; padding: 10px; display: flex; gap: 10px; justify-content: center; align-items: center; ">' +
          `<div style="background: ${w["config"]["colors"][seriesIndex]}; height: 10px; width: 10px; border-radius: 100%;"></div>` +
          `<p>${w["config"]["series"][seriesIndex]["name"]} : ${series[seriesIndex][dataPointIndex]}</p>` +
          "</div>"
        );
      },
      x: {
        show: false,
      },
      y: tooltipDetail,
    },
  };

  const mergedOptions = merge({}, defaultOptions, options);
  return (
    <ApexChart
      series={series}
      type="bar"
      height={310}
      options={mergedOptions}
      {...prop}
    />
  );
};
