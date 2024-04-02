import axios from "axios";
import {config} from "../config";


export interface User {
    id: string,
    name: string,

}

export default class UserService {
    static async get() {
        const response = await axios.get(`${config.API}${config.USERS}/`);
        const users: User[] = response.data;
        return users;
    }
}
