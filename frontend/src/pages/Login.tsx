import React, { useState } from 'react';
import AuthService from './../services/Auth';
import { useNavigate } from 'react-router-dom';
import {Button, Container, Form} from "./styles/Login.styled";

const LoginScreen = () => {
    const navigate = useNavigate();
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [error, setError] = useState<string | null>(null);

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault();
        setError(null);

        try {
            await AuthService.login(email, password);
            console.log("Login success", email, password);
            navigate('/dashboard');
        } catch (error) {
            setError('Wrong password or email');
        }
    };

    return (
        <Container>
            <h2>Login</h2>
            {error && <p style={{ color: 'red' }}>{error}</p>}
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
        </Container>
    );
};

export default LoginScreen;