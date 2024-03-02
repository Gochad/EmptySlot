import axios from "axios";
import {API_URL, LOGIN_PREFIX} from "../config";

export interface BaseEvent {
    event_id: number,
    start: Date
    end: Date,
    title: string,
}
export interface Reservation {
    id: number,
    starttime: string,
    endtime: string,
    confirmed: boolean,
    isreserved: boolean,
}

const merchandisesReqExample: never[] = [
];

const customerReqExample = {
};

export const mapEventToReservationRequests = (event: BaseEvent) => {
    return {
        ID: event.event_id.toString(),
        MerchandisesReq: merchandisesReqExample,
        CustomerReq: customerReqExample,
        Confirmed: false,
        StartTime: event.start.toISOString(),
        EndTime: event.end.toISOString(),
        IsReserved: false,
    };
};

function convertStringToDate(dateString: string) {
    const date = new Date(dateString);
    return date;
}

export const mapReservationToEvent = (reservation: Reservation): BaseEvent => {
    return {
        event_id: Number(reservation.id),
        start: convertStringToDate(reservation.starttime),
        end: convertStringToDate(reservation.endtime),
        title: ""
    };
};
export let EVENTS = []

export class Events {
    static async get() {
        const response = await axios.get(`${API_URL}/reservations/`);
        console.log(response.data);
        return response.data;
    }
}
