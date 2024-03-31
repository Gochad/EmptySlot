import axios from "axios";
import {API_URL, APP_URL, MAIN_PREFIX} from "../config";

export class Reservation {
    static async pay(id: string) {
        const response = await axios.post(`${API_URL}/reservations/${id}/pay?redirect_url=${encodeURIComponent(APP_URL + MAIN_PREFIX)}`)

        // return payment link
        return response.data;
    }
}