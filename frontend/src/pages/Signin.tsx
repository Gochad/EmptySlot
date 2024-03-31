import * as React from 'react';
import Button from '@mui/material/Button';
import CssBaseline from '@mui/material/CssBaseline';
import TextField from '@mui/material/TextField';
import FormControlLabel from '@mui/material/FormControlLabel';
import Checkbox from '@mui/material/Checkbox';
import Link from '@mui/material/Link';
import Grid from '@mui/material/Grid';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import {OAuth, Logo} from "./styles/Login.styled";
import {API_URL} from "../config";
import AuthService from "../services/Auth";
import {errorPopup} from "../services/utils";
import {useNavigate} from "react-router-dom";

const defaultTheme = createTheme();

export default function SignIn() {
    const sso = `${API_URL}/google-sso`;
    const redirectURI = "http://localhost:3001/dashboard";

    const ssoUrl = `${sso}?redirect_uri=${encodeURIComponent(redirectURI)}`;
    const navigate = useNavigate();


    const handleSubmit = async (e: React.FormEvent) => {
        try {
            e.preventDefault();
            const formData = new FormData(e.target as HTMLFormElement);

            const formProps = Object.fromEntries(formData) as { [key: string]: FormDataEntryValue };

            await AuthService.login({
                email: formProps.email.toString(),
                password: formProps.password.toString()
            });
            navigate('/dashboard');
        } catch (error) {
            errorPopup('Wrong password or email');
        }
    };

    return (
        <ThemeProvider theme={defaultTheme}>
            <Container component="main" maxWidth="xs">
                <CssBaseline />
                <Box
                    sx={{
                        marginTop: 8,
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center',
                    }}
                >
                    <Logo src="logo.jpeg" alt="EmptySlot Logo"/>
                    <Typography component="h1" variant="h5">
                        Sign in
                    </Typography>
                    <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
                        <TextField
                            margin="normal"
                            required
                            fullWidth
                            id="email"
                            label="Email Address"
                            name="email"
                            autoComplete="email"
                            autoFocus
                        />
                        <TextField
                            margin="normal"
                            required
                            fullWidth
                            name="password"
                            label="Password"
                            type="password"
                            id="password"
                            autoComplete="current-password"
                        />
                        <FormControlLabel
                            control={<Checkbox value="remember" color="primary" />}
                            label="Remember me"
                        />
                        <Button
                            type="submit"
                            fullWidth
                            variant="contained"
                            sx={{ mt: 3, mb: 2 }}
                        >
                            Sign In
                        </Button>
                        <Grid container>
                            <Grid item xs>
                                <Link href="#" variant="body2">
                                    Forgot password?
                                </Link>
                            </Grid>
                            <Grid item>
                                <Link href="#" variant="body2">
                                    {"Don't have an account? Sign Up"}
                                </Link>
                            </Grid>
                        </Grid>
                    </Box>
                </Box>
                <OAuth>
                    <a href={ssoUrl}>
                        <img src="logo_google.png" alt="Google Logo"/>
                    </a>
                </OAuth>
            </Container>
        </ThemeProvider>
    );
}