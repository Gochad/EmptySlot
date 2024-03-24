import React from "react";
import Navbar from "./../components/Navbar";

import FullCalendar from '../components/FullCalendar';

export default function MainPage() {
    const token = localStorage.getItem('token');
    const isTokenValid = token;

    return isTokenValid ? (
        <div>
            <Navbar />
            <FullCalendar />
        </div>
    ): <div>{`You can't access`}</div>;
}
