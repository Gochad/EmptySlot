import axios from "axios/index";
import {API_URL, LOGIN_PREFIX} from "../config";

interface Event {
    event_id: number,
    title: string,
    start: Date,
    end: Date
}

const merchandisesReqExample: never[] = [
];

const customerReqExample = {
};

const mapEventToReservationRequests = (event) => {
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
export let EVENTS = []

class Events {
    static async get(data: Event) {
        const response = await axios.post(`${API_URL}/reservation`, data);

        return response.data;

    }
}
