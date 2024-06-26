import {config} from "../config";
import axios from "axios";

interface LoginResponse {
    token: string;
    reservation: string;
    email: string;
    role: number;
}

interface LoginData {
    email: string;
    password: string;
}

class AuthService {
    static async login(userData: LoginData){
        const response = await axios.post(`${config.API}${config.LOGIN}`, userData);
        const data: LoginResponse = response.data;
        localStorage.setItem('token', data.token);
        localStorage.setItem('reservation', data.reservation);
        localStorage.setItem('email', data.email);
        localStorage.setItem('role', data.role === 0 ? "admin" : "user");
    }

    static logout() {
        localStorage.removeItem('token');
        localStorage.removeItem('userid');
    }
}

export default AuthService;