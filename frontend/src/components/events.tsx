import axios from "axios";
import {API_URL, LOGIN_PREFIX} from "../config";
import {errorPopup, removeInvalidDates} from "./utils";
import {ProcessedEvent} from "@aldabil/react-scheduler/types";

axios.interceptors.request.use(config => {
    const token = localStorage.getItem('token');
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

export interface BaseEvent {
    event_id: number,
    title: string,
    price: number,
    start: Date
    end: Date,
    title: string,
}
export interface Reservation {
    id: number,
    name: string,
    price: number,
    starttime: string,
    endtime: string,
    confirmed: boolean,
    isreserved: boolean,
}

export const mapEventToReservationRequests = (event: BaseEvent | ProcessedEvent) => {
    if (event && !event.start && !event.end){
        errorPopup("Event obj doesn't have start or end prop")
        return
    }
    return {
        Confirmed: false,
        Name: event.title,
        Price: event.price,
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
        title: reservation.name,
        price: reservation.price,
    };
};

export class Events {
    static async get() {
        const response = await axios.get(`${API_URL}/reservations/`);
        const events: Reservation[] = response.data;
        return removeInvalidDates(events.map(v => mapReservationToEvent(v)));
    }

    static async create(data: ProcessedEvent) {
        const mapped = mapEventToReservationRequests(data);
        await axios.post(`${API_URL}/reservations/`, mapped);
    }

    static async delete(id: string | number) {
        await axios.delete(`${API_URL}/reservations/${id}`);
    }
}
