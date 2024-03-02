import React, {useEffect, useState} from "react";
import {Scheduler} from "@aldabil/react-scheduler";
import {Events, mapReservationToEvent, Reservation, BaseEvent} from "./events";
import {removeInvalidDates} from "./utils";

const Calendar = () => {
    const [events, setEvents] = useState<BaseEvent[]>([]);
    const [error, setError] = useState<string | null>(null);

    const translations = {
        navigation: {
            month: "Month",
                week: "Week",
                day: "Day",
                today: "Today",
            agenda: "Agenda"
        },
        form: {
            addTitle: "Add merchandise",
                editTitle: "Edit merchandise",
                confirm: "Confirm",
                delete: "Delete",
                cancel: "Cancel"
        },
        event: {
            title: "Title",
                start: "Start",
                end: "End",
                allDay: "All Day"
        },
        moreEvents: "More...",
        loading: "Loading..."
    }

    const rerenderEvents = async () => {
        try {
            const events: Reservation[] = await Events.get();
            return removeInvalidDates(events.map(v => mapReservationToEvent(v)))
        } catch (error) {
            setError('problem with loading data');
        }
    };

    useEffect(() => {
        (async () => {
            try {
                const events: BaseEvent[] | undefined = await rerenderEvents();
                if (events) {
                    setEvents(events);
                }
            } catch (error) {
                setError(`Error fetching events: ${error}`);
            }
        })();
    }, []);



    return (
        <div style={{maxWidth: '500px', maxHeight: '100px', display: 'flex', width: '350px'}}>
            <Scheduler
                events={events}
                height={10}
                translations={translations}
                onEventClick={rerenderEvents}
                draggable={true} //admin = true, standard = false
            />
        </div>
    );
}
export default Calendar;