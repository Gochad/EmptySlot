import axios from "axios";
import {config} from "../config";


export interface User {
    id: string,
    name: string,
    username: string,
    email: string;
    address: string,
    phone: string;
    CreatedAt: string;
    UpdatedAt: string;
}

export default class UserService {
    static async get(email: string) {
        const response = await axios.get(`${config.API}${config.USERS}/${email}`);
        console.log(response.data)
        const user: User = response.data;
        return user;
    }
}
