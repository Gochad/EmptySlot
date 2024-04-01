import React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import { List, ListItem } from '@mui/material';
import Navbar from "../components/Navbar";
import Container from "@mui/material/Container";
import Grid from "@mui/material/Grid";

export default function Cart() {
    const items = [
        {
            id: 1,
            name: 'item1',
            description: 'desc1',
            merchandises: ['Merchandiseid1', 'Merchandiseid2'],
            isReserved: false,
        },
        {
            id: 2,
            name: 'item2',
            description: 'desc2',
            merchandises: ['Merchandiseid3', 'Merchandiseid4'],
            isReserved: true,
        },
    ];

    return (
        <>
            <Navbar />
            <Container maxWidth="md">
                <Typography variant="h4" component="h1" gutterBottom>
                    Merchandises in your cart to reserve
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
                                <List>
                                    {item.merchandises.map((merchandise, index) => (
                                        <ListItem key={index}>
                                            {merchandise}
                                        </ListItem>
                                    ))}
                                </List>
                                <Typography variant="body2">
                                    {item.isReserved ? 'Reserved' : 'Available'}
                                </Typography>
                            </CardContent>
                        </Card>
                    ))}
                </Grid>
            </Container>
        </>
    );
}
