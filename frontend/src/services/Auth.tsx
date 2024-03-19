import {API_URL, LOGIN_PREFIX, REGISTER_PREFIX} from "../config";
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
        const response = await axios.post(`${API_URL}${LOGIN_PREFIX}`, userData);

        const data: LoginResponse = response.data;
        localStorage.setItem('token', data.token);
    }

    static logout() {
        localStorage.removeItem('token');
    }
}

export default AuthService;