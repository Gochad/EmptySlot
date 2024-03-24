import React from "react";
import {Btn, Buttons} from "./styles/Navbar.styled";
import {useNavigate} from "react-router-dom";

export default function Navbar() {
    const navigate = useNavigate();

    const logout = (e: React.MouseEvent<HTMLButtonElement>) => {
        e.preventDefault();
        localStorage.removeItem('token');
        navigate('/login');
    };

    return (
        <div>

            <Buttons>
                <Btn>Payments history</Btn>
                <Btn>User</Btn>
                <Btn type="submit"
                     onClick={logout}>Logout
                </Btn>
            </Buttons>

        </div>
    );
}
