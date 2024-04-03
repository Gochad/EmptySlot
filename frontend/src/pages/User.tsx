import React, {useEffect, useState} from "react";
import { Paper, Typography, Grid } from '@mui/material';
import {errorPopup} from "../services/utils";
import {User} from "../services/User";
import UserService from "../services/User";
import Navbar from "../components/Navbar";


//TODO: fixme
export default function UserScreen() {
    const [userData, setUserData] = useState<User>();

    const loadUser = async () => {
        try {
            return await UserService.get(localStorage.getItem('email') as string);
        } catch (error) {
            errorPopup(`error while getting user: ${error}`);
        }
    };

    useEffect(() => {
        loadUser().then(item => {
            if (item) {
                setUserData(item);
            }
        }).catch(error => {
            errorPopup(`error fetching user: ${error}`);
        });
    }, []);


    return (
        <>
            <Navbar />
            <Paper elevation={3} sx={{ padding: 2, margin: 2 }}>
                <Typography variant="h6" gutterBottom>
                    User Details
                </Typography>
                <Grid container spacing={2}>
                    <Grid item xs={12}>
                        <Typography>Email: {userData?.email}</Typography>
                    </Grid>
                    <Grid item xs={12}>
                        <Typography>Username: {userData?.username}</Typography>
                    </Grid>
                    <Grid item xs={12}>
                        <Typography>Address: {userData?.address || 'N/A'}</Typography>
                    </Grid>
                    <Grid item xs={12}>
                        <Typography>Phone: {userData?.phone || 'N/A'}</Typography>
                    </Grid>
                    <Grid item xs={12}>
                        <Typography>Created At: {userData?.CreatedAt}</Typography>
                    </Grid>
                    <Grid item xs={12}>
                        <Typography>Updated At: {userData?.UpdatedAt}</Typography>
                    </Grid>
                </Grid>
            </Paper>
        </>
    );
}
