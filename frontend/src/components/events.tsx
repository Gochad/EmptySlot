import axios from "axios";
import {API_URL, LOGIN_PREFIX} from "../config";

interface Event {
    event_id: number,
    title: string,
    start: Date,
    end: Date
}
interface Reservation {
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

export const mapEventToReservationRequests = (event: Event) => {
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
    if (isNaN(date.getTime())) {
        return null;
    } else {
        return date;
    }
}

export const mapReservationToEvent = (reservation: Reservation) => {
    return {
        event_id: reservation.id,
        start: convertStringToDate(reservation.starttime),
        end: convertStringToDate(reservation.endtime),
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
