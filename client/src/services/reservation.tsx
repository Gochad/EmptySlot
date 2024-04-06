import axios from "axios";
import {config} from "../config";

export interface Reservation {
    id: number,
    name: string,
    description: string,
    isreserved: boolean,
}

export class ReservationService {
    static async pay(id: string) {
        const redirect = config.APP + config.MAIN;
        const url = `${config.API}/${config.RESERVATION}/${id}/pay?redirect_url=${redirect}`
        const response = await axios.get(url)

        // return payment link
        return response.data;
    }

    static async get(id: string) {
        const url = `${config.API}/${config.RESERVATION}/${id}/`
        const response = await axios.get(url)

        return response.data;
    }
}