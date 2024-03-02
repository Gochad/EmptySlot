import React, {useEffect, useState} from "react";
import {Scheduler} from "@aldabil/react-scheduler";
import {Events, mapReservationToEvent} from "./events";

const Calendar = () => {
    const [events, setEvents] = useState([]);
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
            const events = await Events.get();
            return events
        } catch (error) {
            setError('problem with loading data');
        }
    };
    useEffect(() => {
        (async () => {
            try {
                const events = await rerenderEvents();
                if (events && events.length > 0) {
                    const mappedEvent = mapReservationToEvent(events[4]);
                    setEvents([mappedEvent]);
                    console.log(mapReservationToEvent(events[4]));
                }
            } catch (error) {
                console.error("Error fetching events:", error);
            }
        })();
    });



    return (
        <div style={{maxWidth: '500px', maxHeight: '100px', display: 'flex', width: '350px'}}>
            <Scheduler
                events={events}
                height={10}
                translations={translations}
                onEventClick={rerenderEvents}
            />
        </div>
    );
}
export default Calendar;