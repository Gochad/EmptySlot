import React, {useEffect, useState} from 'react';
import AuthService from './../services/Auth';
import { useNavigate } from 'react-router-dom';
import {Button, Container, Form, OAuth} from "./styles/Login.styled";
import {errorPopup} from "../components/utils";
import {API_URL, LOGIN_PREFIX} from "../config";

const LoginScreen = () => {
    const ssoUrl = `${API_URL}/google-sso`;
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
        <Container>
            <h2>Login</h2>
            <Form onSubmit={handleLogin}>
                <div>
                    <label htmlFor="email">Email:</label>
                    <input
                        type="email"
                        id="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                </div>
                <div>
                    <label htmlFor="password">Password:</label>
                    <input
                        type="password"
                        id="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </div>
                <Button type="submit">Login</Button>
            </Form>

            <OAuth>
                <a href={ssoUrl}>
                    <img src="logo_google.png" alt="Google Logo"/>
                </a>
            </OAuth>
        </Container>
    );
};

export default LoginScreen;