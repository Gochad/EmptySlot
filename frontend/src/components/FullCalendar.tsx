import React, {ChangeEvent, useEffect, useState} from "react";
import {Calendar, momentLocalizer, SlotInfo} from "react-big-calendar";
import moment from "moment";
import "react-big-calendar/lib/css/react-big-calendar.css";
import {BaseEvent, EventsService} from "../services/events";
import {errorPopup, successPopup} from "../services/utils";
import Modal from 'react-modal';
import AddEventModal from "./AddEventModal";
import ShowEventModal from "./ShowEventModal";
import "./styles/calendar.css";
import {CalendarStyles} from "./styles/FullCalendar.styled";

Modal.setAppElement('#root');

moment.locale("en-GB");
const localizer = momentLocalizer(moment);

export default function FullCalendar() {
    const [events, setEvents] = useState<BaseEvent[]>([]);
    const [newEvent, setNewEvent] = useState<BaseEvent>({ event_id: 0, title: '', description: '', category_id: '', price: 0, start: new Date(), end: new Date() });
    const [currentEventId, setCurrentEventId] = useState<string>('');

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

    const handleSelect = ({ start, end }: SlotInfo) => {
        setNewEvent({ event_id: 0, title: '', description: '', category_id: '', price: 0, start: start, end: end });
        openModal('addEventModal');

    };
    const handleSave = async (categoryId: string) => {
        try {
            const updatedEvent = {...newEvent, category_id: categoryId};

            if (newEvent.title) {
                setEvents([
                    ...events,
                    updatedEvent
                ]);
            }
            await EventsService.create(updatedEvent);
            successPopup(`event added`);
        } catch (error) {
            errorPopup(`error while saving new event: ${error}`);
        } finally {
            closeModal('addEventModal');
        }
    };
    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;

        setNewEvent(prevEvent => ({
            ...prevEvent,
            [name]: name === 'price' ? Number(value) : value
        }));
    };

    const handleAssignToReservation = async() => {
        try {
            const e = await EventsService.getById(currentEventId);
            e.reservation_id = localStorage.getItem('reservation') as string
            await EventsService.update(e)
            successPopup(`event assign to current reservation`);
        } catch (error) {
            console.log(error)
            errorPopup(`error while assign current reservation: ${error}`);
        } finally {
            closeModal('showEventModal');
        }
    };

    const handleEventSelect = (event: BaseEvent) => {
        setCurrentEventId(String(event.event_id));
        openModal('showEventModal');
    };

    const rerenderEvents = async () => {
        try {
            const events: BaseEvent[] = await EventsService.get();
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
        <div>
            <Calendar
                localizer={localizer}
                events={events}
                startAccessor="start"
                endAccessor="end"
                style={CalendarStyles}
                selectable
                onSelectSlot={handleSelect}
                defaultView="week"
                views={["week","agenda"]}
                onSelectEvent={handleEventSelect}
                min={new Date(0, 0, 0, 6, 0, 0)}
                max={new Date(0, 0, 0, 22, 0, 0)}
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
                handleAssign={handleAssignToReservation}
            />
        </div>
    );
}