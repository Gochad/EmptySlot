import React from "react";
import Calendar from "./../components/Calendar";
import Navbar from "./../components/Navbar";

const MainPage = () => {
    const token = localStorage.getItem('token');
    const isTokenValid = token;

    return isTokenValid ? (
        <div>
            <Navbar />
            <div>
                <Calendar />
            </div>
        </div>
    ): <div>You can't access</div>;
}
export default MainPage;