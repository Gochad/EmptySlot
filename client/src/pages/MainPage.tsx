import React from "react";
import Navbar from "./../components/Navbar";
import FullCalendar from '../components/FullCalendar';
import {CalendarHeader} from "./styles/MainPage.styled";

export default function MainPage() {
    const token = localStorage.getItem('token');
    const isTokenValid = token;

    return isTokenValid ? (
        <div>
            <Navbar cart/>

            <CalendarHeader>
                Calendar
            </CalendarHeader>

            <FullCalendar />
        </div>
    ): <div>{`You can't access`}</div>;
}
