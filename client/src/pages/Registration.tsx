import React from 'react';
import { useNavigate } from 'react-router-dom';
import Container from '@mui/material/Container';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';
import Button from '@mui/material/Button';
import RegisterService from "../services/Register";
import {errorPopup} from "../services/utils";
import { Box } from '@mui/material';

export default function RegistrationScreen() {
    const navigate = useNavigate();

    const queryParams = new URLSearchParams(location.search);
    const usertype = queryParams.get('usertype') || "user";
    const handleRegister = async (e: React.FormEvent) => {
        try {
            e.preventDefault();
            const formData = new FormData(e.target as HTMLFormElement);
            const formProps = Object.fromEntries(formData) as { [key: string]: FormDataEntryValue };

            await RegisterService.register({
                username: formProps.username.toString(),
                password: formProps.password.toString(),
                email: formProps.email.toString()
            }, usertype);
            navigate('/login');
        } catch (error) {
            errorPopup('Wrong password or email');
        }
    };

    return (
        <Container maxWidth="sm" sx={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            justifyContent: 'center',
            height: '100vh'
        }}>
            <Typography variant="h5" component="h2" sx={{ mb: 2 }}>
                Register
            </Typography>
            <Box component="form" onSubmit={handleRegister} noValidate sx={{ mt: 1 }}>
                <TextField
                    label="Username"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    name="username"
                    required
                />
                <TextField
                    label="Password"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    name="password"
                    type="password"
                    required
                />
                <TextField
                    label="Email"
                    variant="outlined"
                    fullWidth
                    margin="normal"
                    name="email"
                    type="email"
                    required
                />
                <Box sx={{ display: 'flex', justifyContent: 'center', mt: 3, mb: 2 }}>
                    <Button
                        type="submit"
                        variant="contained"
                    >
                        Register
                    </Button>
                </Box>
            </Box>
        </Container>
    );
}
