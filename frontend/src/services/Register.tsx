import axios from "axios";

class RegisterService {
    private static apiUrl: string = 'http://localhost:8080';

    static async register(userData: any){
        return await axios.post(`${this.apiUrl}/register`, userData);
    }
}

export default RegisterService;