import React from "react";
import Navbar from "./../components/Navbar";

import FullCalendar from '../components/FullCalendar';
import { Container } from '@mui/material';

const MainPage = () => {
    const token = localStorage.getItem('token');
    const isTokenValid = token;

    return isTokenValid ? (
        <div>
            <Navbar />
            {/*<div>*/}
            {/*    <Calendar />*/}
            {/*</div>*/}
            <FullCalendar />
        </div>
    ): <div>You can't access</div>;
}
export default MainPage;