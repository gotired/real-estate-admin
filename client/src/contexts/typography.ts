import type { TypographyVariantsOptions } from "@mui/material/styles";
import { Manrope } from "next/font/google";

const manrope = Manrope({ subsets: ["latin"] });

export const typography: TypographyVariantsOptions = {
  fontFamily: [
    manrope.style.fontFamily,
    "Montserrat",
    "-apple-system",
    "BlinkMacSystemFont",
    '"Segoe UI"',
    "Roboto",
    '"Helvetica Neue"',
    "Arial",
    "sans-serif",
    '"Apple Color Emoji"',
    '"Segoe UI Emoji"',
    '"Segoe UI Symbol"',
  ].join(","),
};
