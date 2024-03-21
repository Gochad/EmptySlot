import axios from "axios";
import {API_URL} from "../config";

export class Reservation {
    static async pay(id: string) {
        const response = await axios.post(`${API_URL}/reservations/${id}/pay`);
        const link = response.data;
        console.log("PAYMENT LINK", link);
        return link;
    }
}