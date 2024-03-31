import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LandingPage from "./pages/LandingPage";
import MainPage from "./pages/MainPage";
import SignIn from "./pages/Signin";
import RegistrationScreen from "./pages/Registration";
import {MAIN_PREFIX, LOGIN_PREFIX, REGISTER_PREFIX} from "./config";

function App() {
    return (
        <Router>
            <div>
                <Routes>
                    <Route path="" element={<LandingPage />} />
                    <Route path={MAIN_PREFIX} element={<MainPage />} />
                    <Route path={LOGIN_PREFIX} element={<SignIn />} />
                    <Route path={REGISTER_PREFIX} element={<RegistrationScreen />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
