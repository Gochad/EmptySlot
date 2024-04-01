import axios from "axios";
import {config} from "../config";

export interface Reservation {
    id: number,
    name: string,
    description: string,
    merchandises: string[],
    isreserved: boolean,
}

export class Reservation {
    static async pay(id: string) {
        const redirect = `${encodeURIComponent(config.APP + config.MAIN)}`
        const url = `${config.API}/${config.RESERVATION}/${id}/pay?redirect_url=${redirect}`
        const response = await axios.post(url)

        // return payment link
        return response.data;
    }
}