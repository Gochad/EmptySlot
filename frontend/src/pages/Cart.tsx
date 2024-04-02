import React, {useEffect, useState} from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import Navbar from "../components/Navbar";
import Container from "@mui/material/Container";
import Grid from "@mui/material/Grid";
import {errorPopup} from "../services/utils";
import {EventsService, Merchandise} from "../services/events";
import {Reservation, ReservationService} from "../services/reservation";

export default function Cart() {
    const [merchandises, setMerchandises] = useState<Merchandise[]>([]);

    const createPaymentLink = async () => {
        try {
            const link = await ReservationService.pay(localStorage.getItem('reservation') as string);
            window.location.href = link;
        } catch (error) {
            errorPopup(`problem with creating payment link: ${error}`);
        }
    };

    const loadReservation = async () => {
        try {
            const reservationID = localStorage.getItem('reservation') as string;
            const items: Merchandise[] = await EventsService.getByReservation(reservationID);
            return items;
        } catch (error) {
            errorPopup(`error while getting all reservations: ${error}`);
        }
    };

    useEffect(() => {
        loadReservation().then(items => {
            if (items) {
                setMerchandises(items);
            }
        }).catch(error => {
            errorPopup(`error fetching categories: ${error}`);
        });
    }, []);

    return (
        <>
            <Navbar />
            <Container maxWidth="md">
                <Typography variant="h4" component="h1" gutterBottom>
                    Merchandises to reserve
                </Typography>
                <button type="submit" className="btn btn-primary" onClick={createPaymentLink}>Pay</button>
                <Grid container spacing={2}>
                    {merchandises.map((item) => (
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
