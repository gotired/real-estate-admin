interface BarChartProps {
  colors: Array<string>;
  legend: Array<string>;
  ylabel: string;
  tooltipDetail: { formatter?(val: number, opts?: unknown): string };
}

interface BarChartComponentProp {
  series: Array<SeriesProps>;
  colors: Array<string>;
  legend: Array<string>;
  ylabel: string;
  tooltipDetail: { formatter?(val: number, opts?: unknown): string };
  [key: string]: unknown;
}

interface SeriesProps {
  name?: string;
  data: Array<number | null>;
}
