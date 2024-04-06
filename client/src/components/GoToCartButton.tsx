import React from 'react';
import Button from '@mui/material/Button';
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import { useNavigate } from 'react-router-dom';

export default function GoToCartButton() {
    const navigate = useNavigate();

    const goToCart = () => {
        navigate('/cart');
    };

    return (
        <Button variant="contained" color="secondary" onClick={goToCart} startIcon={<ShoppingCartIcon />}>
            Cart
        </Button>
    );
}
