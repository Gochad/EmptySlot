import React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import Navbar from "../components/Navbar";
import Container from "@mui/material/Container";
import Grid from "@mui/material/Grid";
import {errorPopup, successPopup} from "../services/utils";
import {ReservationService} from "../services/reservation";

export default function Cart() {
    const items = [
        {
            id: 1,
            name: 'item1',
            description: 'desc1',
            price: 1,
            confirmed: false,
        },
        {
            id: 2,
            name: 'item2',
            description: 'desc2',
            price: 1,
            confirmed: false,
        },
    ];

    const loadReservation = async () => {
        try {
            const reservationID = localStorage.getItem('reservation') as string;
            await ReservationService.get(reservationID);
            successPopup(`event added`);
        } catch (error) {
            errorPopup(`error while getting all reservations: ${error}`);
        }
    };

    return (
        <>
            <Navbar />
            <Container maxWidth="md">
                <Typography variant="h4" component="h1" gutterBottom>
                    Merchandises to reserve
                </Typography>
                <Grid container spacing={2}>
                    {items.map((item) => (
                        <Card key={item.id} sx={{ marginBottom: 2 }}>
                            <CardContent>
                                <Typography variant="h5" component="div">
                                    {item.name}
                                </Typography>
                                <Typography sx={{ mb: 1.5 }} color="text.secondary">
                                    {item.description}
                                </Typography>
                                <Typography sx={{ mb: 1.5 }} color="text.secondary">
                                    {item.price}
                                </Typography>
                                <Typography variant="body2">
                                    {item.confirmed ? 'Reserved' : 'Available'}
                                </Typography>
                            </CardContent>
                        </Card>
                    ))}
                </Grid>
            </Container>
        </>
    );
}
