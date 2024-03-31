import React from 'react';
import { AppBar, Toolbar, Typography, Button, Container, Box } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import {Logo} from "./styles/LandingPage.styled";

export default function LandingPage(){
    const navigate = useNavigate();

    return (
        <>
            <AppBar position="static" color="primary" elevation={0}>
                <Toolbar>
                    <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }}>
                        <Logo src="logo.jpeg" alt="EmptySlot Logo" />
                        <Typography variant="h6" color="inherit" noWrap>
                            EmptySlot
                        </Typography>
                    </Box>
                    <Box sx={{ flexGrow: 1 }} />
                    <Button color="inherit" onClick={() => navigate('/login')}>
                        Log in
                    </Button>
                    <Button color="inherit" onClick={() => navigate('/register')}>
                        Register
                    </Button>
                </Toolbar>
            </AppBar>

            <Container maxWidth="lg" component="main" sx={{ mt: 20, mb: 6 }}>
                <Typography variant="h2" align="center" color="textPrimary" gutterBottom>
                    Welcome to the EmptySlot app! This app provides easy payments for calendar-based services!
                </Typography>
                <Typography variant="h5" align="center" color="textSecondary" sx={{ mt: 3 }}>
                    You can create either a user or an admin account. As an admin, you can create your own service calendar and add events to it.
                    As a user, you can select services, make reservations, and pay for them.
                </Typography>
            </Container>
        </>
    );
}
