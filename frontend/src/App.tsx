import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LandingPage from "./pages/LandingPage";
import MainPage from "./pages/MainPage";
import SignIn from "./pages/Signin";
import RegistrationScreen from "./pages/Registration";
import CategoriesScreen from "./pages/Categories";
import {config} from "./config";

function App() {
    return (
        <Router>
            <div>
                <Routes>
                    <Route path="" element={<LandingPage />} />
                    <Route path={config.MAIN_PREFIX} element={<MainPage />} />
                    <Route path={config.LOGIN_PREFIX} element={<SignIn />} />
                    <Route path={config.REGISTER_PREFIX} element={<RegistrationScreen />} />
                    <Route path={config.CATEGORIES_PREFIX} element={<CategoriesScreen />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
