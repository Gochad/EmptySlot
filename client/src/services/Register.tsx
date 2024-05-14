import axios from "axios";
import {config} from "../config";

interface RegisterData {
    username: string;
    email: string;
    password: string;
}

class RegisterService {
    static async register(userData: RegisterData, userType: string){
        if (userType === "admin") {
            return await axios.post(`${config.API}${config.REGISTER}?usertype=admin`, userData);
        }
        return await axios.post(`${config.API}${config.REGISTER}?usertype=user`, userData);
    }
}

export default RegisterService;