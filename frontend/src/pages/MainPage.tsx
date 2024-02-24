import React from "react";
import Calendar from "./../components/Calendar";
import {useNavigate} from "react-router-dom";
import {LogoutBtn} from "./styles/MainPage";

const MainPage = () => {
    const navigate = useNavigate()
    const handleChange = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        navigate('/login');
    };

    return (
        <div>
            <div>
                <Calendar />
            </div>

            <LogoutBtn type="submit"
                    onClick={handleChange}>Logout
            </LogoutBtn>
        </div>

    );
}
export default MainPage;