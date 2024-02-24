import axios from "axios";
import {API_URL, REGISTER_PREFIX} from "../config";

interface RegisterData {
    username: string;
    email: string;
    password: string;
}

class RegisterService {
    static async register(userData: RegisterData){
        return await axios.post(`${API_URL}${REGISTER_PREFIX}`, userData);
    }
}

export default RegisterService;