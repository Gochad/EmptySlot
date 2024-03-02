import {BaseEvent} from "./events";

function isInvalidDate(date: Date) {
    return isNaN(date.getTime());
}

export function removeInvalidDates(events: BaseEvent[]) {
    return events.filter(v => !isInvalidDate(v.start) && !isInvalidDate(v.start));
}