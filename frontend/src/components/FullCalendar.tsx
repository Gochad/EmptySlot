import React, {useEffect, useState} from "react";
import { Calendar, momentLocalizer } from "react-big-calendar";
import moment from "moment";
import "react-big-calendar/lib/css/react-big-calendar.css";
import {Events, BaseEvent} from "./events";
import {errorPopup, successPopup} from "./utils";
import Modal from 'react-modal';
import AddEventModal from "./AddEventModal";
import ShowEventModal from "./ShowEventModal";

Modal.setAppElement('#root');

moment.locale("en-GB");
const localizer = momentLocalizer(moment);

export default function FullCalendar() {
    const [events, setEvents] = useState<BaseEvent[]>([]);
    const [newEvent, setNewEvent] = useState({ start: null, end: null, title: '', price: 0 });
    const [currentEventId, setCurrentEventId] = useState<string>();

    const [modals, setModals] = useState({
        addEventModal: false,
        showEventModal: false,
    });

    const openModal = (modalName: string) => {
        setModals({ ...modals, [modalName]: true });
    };

    const closeModal = (modalName: string) => {
        setModals({ ...modals, [modalName]: false });
    };

    const handleSelect = ({ start, end }) => {
        setNewEvent({ start, end, title: '' });
        openModal('addEventModal');

    };
    const handleSave = async () => {
        if (newEvent.title) {
            setEvents([
                ...events,
                newEvent
            ]);
        }

        try {
            await Events.create(newEvent);
            successPopup(`event added`);
        } catch (error) {
            errorPopup(`error while saving new event: ${error}`);
        } finally {
            closeModal('addEventModal');
        }
    };
    const handleChange = (e) => {
        const { name, value } = e.target;

        setNewEvent(prevEvent => ({
            ...prevEvent,
            [name]: name === 'price' ? Number(value) : value
        }));
    };

    const handleEventSelect = (event: BaseEvent) => {
        console.log(event)
        setCurrentEventId(event.event_id);
        openModal('showEventModal');
    };

    const rerenderEvents = async () => {
        try {
            const events: BaseEvent[] = await Events.get();
            return events;
        } catch (error) {
            errorPopup(`problem with loading data: ${error}`);
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
        <div className="App">
            <Calendar
                localizer={localizer}
                events={events}
                startAccessor="start"
                endAccessor="end"
                style={{ height: 500 }}
                selectable
                onSelectSlot={handleSelect}
                defaultView="week"
                views={["week","agenda"]}
                onSelectEvent={handleEventSelect}
            />
            <AddEventModal
                modalIsOpen={modals.addEventModal}
                handleCloseModal={() => closeModal("addEventModal")}
                handleSave={handleSave}
                event={newEvent}
                handleChange={handleChange}
            />
            <ShowEventModal
                modalIsOpen={modals.showEventModal}
                handleCloseModal={() => closeModal("showEventModal")}
                eventId={currentEventId}
            />
        </div>
    );
}