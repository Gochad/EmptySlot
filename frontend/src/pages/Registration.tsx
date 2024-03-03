import {useNavigate} from "react-router-dom";
import React, {ChangeEvent, useState} from "react";
import RegisterService from "../services/Register";
import {errorPopup} from "../components/utils";
import {Button, Container, Form} from "./styles/Registration.styled";


const RegistrationScreen = () => {
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        username: '',
        password: '',
        email: ''
    });

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        setFormData({ ...formData, [e.target.name]: e.target.value });
    };

    const handleRegister = async (e: React.FormEvent) => {
        e.preventDefault();

        try {
            await RegisterService.register(formData);
            navigate('/login');
        } catch (error) {
            errorPopup('Wrong password or email');
        }
    };


    return (
        <Container>
            <h2>Register</h2>
            <Form onSubmit={handleRegister}>
                <div>
                    <label htmlFor="username">Username:</label>
                    <input
                        type="text"
                        id="username"
                        name="username"
                        value={formData.username}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label htmlFor="password">Password:</label>
                    <input
                        type="password"
                        id="password"
                        name="password"
                        value={formData.password}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div>
                    <label htmlFor="email">Email:</label>
                    <input
                        type="email"
                        id="email"
                        name="email"
                        value={formData.email}
                        onChange={handleChange}
                        required
                    />
                </div>
                <Button type="submit">Register</Button>
            </Form>
        </Container>

    );
};

export default RegistrationScreen;