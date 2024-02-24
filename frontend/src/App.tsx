import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import LoginScreen from "./pages/Login";
import MainPage from "./pages/MainPage";


function App() {
    return (
        <Router>
            <div>
                <Routes>
                    <Route path="/login" element={<LoginScreen />} />
                    <Route path="/dashboard" element={<MainPage />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
