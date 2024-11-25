import React, {
  PropsWithChildren,
  createContext,
  useEffect,
  useState,
} from "react";
import { ThemeProvider } from "@mui/material/styles";
import { DarkTheme, LightTheme } from "./theme";

type ColorModeContextType = {
  mode: string;
  setMode: () => void;
};

export const ColorModeContext = createContext<ColorModeContextType>(
  {} as ColorModeContextType
);

export const ColorModeContextProvider: React.FC<PropsWithChildren> = ({
  children,
}) => {
  const [mode, setMode] = useState<string>("light");

  useEffect(() => {
    if (typeof window !== "undefined") {
      const savedMode = window.localStorage.getItem("color-mode");
      if (savedMode) {
        setMode(savedMode);
      } else {
        const isSystemPreferenceDark = window.matchMedia(
          "(prefers-color-scheme: dark)"
        ).matches;
        setMode(isSystemPreferenceDark ? "dark" : "light");
      }
    }
  }, []);

  useEffect(() => {
    if (typeof window !== "undefined") {
      window.localStorage.setItem("color-mode", mode);
    }
  }, [mode]);

  const setColorMode = () => {
    setMode((prevMode) => (prevMode === "light" ? "dark" : "light"));
  };

  return (
    <ColorModeContext.Provider value={{ setMode: setColorMode, mode }}>
      <ThemeProvider theme={mode === "light" ? LightTheme : DarkTheme}>
        {children}
      </ThemeProvider>
    </ColorModeContext.Provider>
  );
};
