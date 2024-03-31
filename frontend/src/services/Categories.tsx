import axios from "axios/index";
import {config} from "../config";


interface CategoryData {
    username: string;
    email: string;
    password: string;
}

class RegisterService {
    static async categories(data: CategoryData){
        return await axios.post(`${config.API}${config.REGISTER}`, data);
    }
}

export default RegisterService;