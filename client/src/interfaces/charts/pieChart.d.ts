interface PieChartProps {
  series: Array<number>;
  colors: Array<string>;
  [key: string]: unknown;
}

interface PieChartCardProps extends PieChartProps {
  title: string;
  value: number;
  [key: string]: unknown;
}
