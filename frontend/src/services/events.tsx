import axios from "axios";
import {config} from "../config";
import {ProcessedEvent} from "@aldabil/react-scheduler/types";
import {errorPopup, removeInvalidDates} from "./utils";

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
    description: string,
    price: number,
    start: Date
    end: Date,
}
export interface Merchandise {
    id: number,
    name: string,
    description: string,
    price: number,
    starttime: string,
    endtime: string,
    confirmed: boolean,
}

export const mapEventToMerchandiseRequests = (event: BaseEvent | ProcessedEvent) => {
    if (event && !event.start && !event.end){
        errorPopup("Event obj doesn't have start or end prop")
        return
    }
    return {
        Confirmed: false,
        Name: event.title,
        Description: event.description,
        Price: event.price,
        StartTime: event.start.toISOString(),
        EndTime: event.end.toISOString(),
    };
};

function convertStringToDate(dateString: string) {
    const date = new Date(dateString);
    return date;
}

export const mapReservationToEvent = (merch: Merchandise): BaseEvent => {
    return {
        event_id: Number(merch.id),
        start: convertStringToDate(merch.starttime),
        end: convertStringToDate(merch.endtime),
        title: merch.name,
        description: merch.description,
        price: merch.price,
    };
};

export class Events {
    static async get() {
        const response = await axios.get(`${config.API}${config.MERCH}/`);
        const events: Merchandise[] = response.data;
        return removeInvalidDates(events.map(v => mapReservationToEvent(v)));
    }

    static async create(data: ProcessedEvent) {
        const mapped = mapEventToMerchandiseRequests(data);
        await axios.post(`${config.API}${config.MERCH}/`, mapped);
    }

    static async delete(id: string | number) {
        await axios.delete(`${config.API}${config.MERCH}/${id}`);
    }
}
