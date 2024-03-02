import {BaseEvent} from "./events";
import { toast } from 'react-toastify';

function isInvalidDate(date: Date) {
    return isNaN(date.getTime());
}

export function removeInvalidDates(events: BaseEvent[]) {
    return events.filter(v => !isInvalidDate(v.start) && !isInvalidDate(v.start));
}

export function makeErrorPopup(msg: string) {
    toast.error(msg, {
        position: "bottom-center",
        autoClose: 5000,
    })
}