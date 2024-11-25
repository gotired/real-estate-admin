import { ColorModeContext } from "@/contexts";
import {
  Box,
  Button,
  Checkbox,
  FormControlLabel,
  InputBase,
  Link,
  TextField,
  useTheme,
} from "@mui/material";
import Container from "@mui/material/Container";
import Typography from "@mui/material/Typography";
import Image from "next/image";
import pic from "@/assets/login.jpeg";
import CustomButton from "@/components/common/CustomButton";

export default function Login() {
  const theme = useTheme();

  return (
    <Box
      sx={{
        display: "grid",
        gridTemplateColumns: "repeat(2, 1fr)",
        height: "100vh",
        alignItems: "center",
      }}
    >
      <Container maxWidth="xs">
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
            gap: "34px",
          }}
        >
          <Box display="flex" flexDirection="column" gap="10px">
            <Typography
              variant="h3"
              fontWeight="700"
              color={theme.palette.text.primary}
            >
              Welcome back
            </Typography>
            <Typography fontWeight="400" color={theme.palette.text.secondary}>
              Welcome back! Please enter your details.
            </Typography>
          </Box>
          <Box
            display="flex"
            flexDirection="column"
            justifyContent="center"
            gap="15px"
          >
            <Box display="flex" flexDirection="column" gap="15px">
              <Box display="flex" flexDirection="column" gap="5px">
                <Typography variant="subtitle2" fontWeight="500">
                  Email
                </Typography>
                <InputBase
                  sx={{
                    height: "46px",
                    border: 1,
                    borderRadius: "10px",
                    padding: "11px 12px",
                    color: theme.palette.text.primary,
                    borderColor: theme.palette.divider,
                    "&.Mui-focused": {
                      borderColor: theme.palette.text.secondary,
                    },
                  }}
                  type="email"
                  placeholder="Enter your email"
                  fullWidth
                />
              </Box>
              <Box display="flex" flexDirection="column" gap="5px">
                <Typography variant="subtitle2" fontWeight="500">
                  Password
                </Typography>
                <InputBase
                  sx={{
                    height: "46px",
                    border: 1,
                    borderRadius: "10px",
                    padding: "11px 12px",
                    color: theme.palette.text.primary,
                    borderColor: theme.palette.divider,
                    "&.Mui-focused": {
                      borderColor: theme.palette.text.secondary,
                    },
                  }}
                  type="password"
                  placeholder="**********"
                  fullWidth
                />
              </Box>
            </Box>
            <Box
              display="flex"
              justifyContent="space-between"
              alignItems="center"
            >
              <Box display="flex" alignItems="center">
                <Checkbox disableRipple size="small" color="info" />
                <Typography>Remember</Typography>
              </Box>
              <Link href="/reset" color={theme.palette.primary.contrastText}>
                <Typography
                  variant="subtitle2"
                  color={theme.palette.primary.contrastText}
                >
                  Forgot Password?
                </Typography>
              </Link>
            </Box>
            <CustomButton
              color={theme.palette.text.primary}
              backgroundColor={theme.palette.primary.contrastText}
              title="Login"
              fullWidth
            />
            <Box display="flex" justifyContent="center" gap="8px">
              <Typography
                variant="subtitle2"
                color={theme.palette.text.secondary}
              >
                Don’t have an account?
              </Typography>
              <Link href="/register" color={theme.palette.primary.contrastText}>
                <Typography
                  variant="subtitle2"
                  color={theme.palette.primary.contrastText}
                >
                  Sign up
                </Typography>
              </Link>
            </Box>
          </Box>
        </Box>
      </Container>
      <Box>
        <Image
          src={pic}
          style={{
            height: "100vh",
          }}
          alt="Picture of the author"
        />
      </Box>
    </Box>
  );
}
