import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LandingPage from "./pages/LandingPage";
import MainPage from "./pages/MainPage";
import SignIn from "./pages/Signin";
import RegistrationScreen from "./pages/Registration";
import CategoriesScreen from "./pages/Categories";
import Cart from "./pages/Cart";
import {config} from "./config";

function App() {
    return (
        <Router>
            <div>
                <Routes>
                    <Route path="" element={<LandingPage />} />
                    <Route path={config.MAIN} element={<MainPage />} />
                    <Route path={config.LOGIN} element={<SignIn />} />
                    <Route path={config.REGISTER} element={<RegistrationScreen />} />
                    <Route path={config.CATEGORIES} element={<CategoriesScreen />} />
                    <Route path={config.CART} element={<Cart />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
