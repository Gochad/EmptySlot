import axios from "axios";
import {config} from "../config";

interface RegisterData {
    username: string;
    email: string;
    password: string;
}

class RegisterService {
    static async register(userData: RegisterData){
        return await axios.post(`${config.API}${config.REGISTER}`, userData);
    }
}

export default RegisterService;