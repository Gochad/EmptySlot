import React, {useState} from 'react';
import AuthService from './../services/Auth';
import { useNavigate } from 'react-router-dom';
import {OAuth, Logo} from "./styles/Login.styled";
import {errorPopup} from "../components/utils";
import {API_URL} from "../config";
import { Container, TextField, Button, Typography } from '@mui/material';

const LoginScreen = () => {
    const sso = `${API_URL}/google-sso`;
    const redirectURI = "http://localhost:3001/dashboard";

    const ssoUrl = `${sso}?redirect_uri=${encodeURIComponent(redirectURI)}`;
    const navigate = useNavigate();
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();

        try {
            await AuthService.login({email, password});
            navigate('/dashboard');
        } catch (error) {
            errorPopup('Wrong password or email');
        }
    };

    return (
        <Container component="main" maxWidth="xs">
            <Logo src="logo.jpeg" alt="EmptySlot Logo"/>
            <div style={{ marginTop: 8, display: 'flex', flexDirection: 'column', alignItems: 'center', }}>
                <Typography component="h1" variant="h5">
                    Logowanie
                </Typography>
                <form style={{ width: '100%', marginTop: 3 }} onSubmit={handleLogin}>
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        id="email"
                        label="Adres Email"
                        name="email"
                        autoComplete="email"
                        autoFocus
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        InputLabelProps={{
                            shrink: true,
                        }}
                    />
                    <TextField
                        variant="outlined"
                        margin="normal"
                        required
                        fullWidth
                        name="password"
                        label="Hasło"
                        type="password"
                        id="password"
                        autoComplete="current-password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        InputLabelProps={{
                            shrink: true,
                        }}
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        color="primary"
                        style={{ margin: '24px 0px 16px' }}
                    >
                        Zaloguj się
                    </Button>
                </form>
            </div>
            <OAuth>
                <a href={ssoUrl}>
                    <img src="logo_google.png" alt="Google Logo"/>
                </a>
            </OAuth>
        </Container>
    );
};

export default LoginScreen;