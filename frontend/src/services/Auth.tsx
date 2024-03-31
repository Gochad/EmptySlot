import {config} from "../config";
import axios from "axios";

interface LoginResponse {
    token: string;
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
    }

    static logout() {
        localStorage.removeItem('token');
    }
}

export default AuthService;