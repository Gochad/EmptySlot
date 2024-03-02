import React, {useEffect, useState} from "react";
import {Scheduler} from "@aldabil/react-scheduler";
import {Events, BaseEvent} from "./events";
import {translations} from "./translations";
import {errorPopup, successPopup} from "./utils";
import {EventActions, ProcessedEvent} from "@aldabil/react-scheduler/types";

const Calendar = () => {
    const [events, setEvents] = useState<BaseEvent[]>([]);

    const rerenderEvents = async () => {
        try {
            const events: BaseEvent[] = await Events.get();
            return events;
        } catch (error) {
            errorPopup(`problem with loading data: ${error}`);
        }
    };

    const onConfirm = async (event: ProcessedEvent, action: EventActions): Promise<any> => {
        try {
            switch (action) {
                case "create":
                    await Events.create(event);
                    successPopup(`Event added`);
                    return events;
                    break;
                default:
                    errorPopup(`Unhandled action: ${action}`);
                    return;
            }
        } catch (error) {
            errorPopup(`Error while confirmation: ${error}`);
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
                errorPopup(`error fetching events: ${error}`);
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
                onConfirm={onConfirm}
            />
        </div>
    );
}
export default Calendar;