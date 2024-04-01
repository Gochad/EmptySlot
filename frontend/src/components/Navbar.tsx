import React from "react";
import {useNavigate} from "react-router-dom";
import {AppBar, Box, Button, ButtonBase, Toolbar, Typography} from "@mui/material";
import {Logo} from "./styles/logo.styled";
import { config } from "../config";
import GoToCartButton from "./GoToCartButton";

interface NavbarProps {
    cart?: boolean;
}

export default function Navbar({cart}: NavbarProps) {
    const navigate = useNavigate();

    const logout = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        localStorage.removeItem('token');
        navigate('/login');
    };

    const categories = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        navigate(config.CATEGORIES);
    };

    const logo = () => {
        navigate(config.MAIN);
    };

    return (
        <AppBar position="static" color="primary" elevation={0}>
            <Toolbar>
                <ButtonBase onClick={logo}>
                    <Box sx={{ display: 'flex', alignItems: 'center', gap: 2 }} >
                        <Logo src="logo.jpeg" alt="EmptySlot Logo"/>
                        <Typography variant="h6" color="inherit" noWrap>
                            EmptySlot
                        </Typography>
                    </Box>
                </ButtonBase>
                <Box sx={{ flexGrow: 1 }} />
                {cart ? <GoToCartButton /> : null}
                <Button color="inherit" onClick={categories}>
                    Service categories
                </Button>
                <Button color="inherit" onClick={() => console.log("user panel")}>
                    User
                </Button>
                <Button color="inherit" onClick={() => console.log("history payments")}>
                    Payments history
                </Button>
                <Button color="inherit" onClick={logout}>
                    Logout
                </Button>
            </Toolbar>
        </AppBar>
    );
}
