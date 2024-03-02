import React, {useEffect, useState} from "react";
import {Scheduler} from "@aldabil/react-scheduler";
import {Events, BaseEvent} from "./events";
import {translations} from "./translations";
import {makeErrorPopup} from "./utils";

const Calendar = () => {
    const [events, setEvents] = useState<BaseEvent[]>([]);

    const rerenderEvents = async () => {
        try {
            const events: BaseEvent[] = await Events.get();
            return events
        } catch (error) {
            makeErrorPopup(`problem with loading data: ${error}`);
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
                makeErrorPopup(`error fetching events: ${error}`);
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