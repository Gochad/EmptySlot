import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LoginScreen from "./pages/Login";
import MainPage from "./pages/MainPage";
import RegistrationScreen from "./pages/Registration";
import {MAIN_PREFIX, LOGIN_PREFIX, REGISTER_PREFIX} from "./config";


function App() {
    return (
        <Router>
            <div>
                <Routes>
                    <Route path={LOGIN_PREFIX} element={<LoginScreen />} />
                    <Route path={REGISTER_PREFIX} element={<RegistrationScreen />} />
                    <Route path={MAIN_PREFIX} element={<MainPage />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
