import React, {useEffect, useState} from "react";
import { Calendar, momentLocalizer } from "react-big-calendar";
import moment from "moment";
import "react-big-calendar/lib/css/react-big-calendar.css";
import {Events} from "./events";
import {errorPopup, successPopup} from "./utils";
import Modal from 'react-modal';
import {Form, modalStyles} from "./styles/FullCalendar.styled";

Modal.setAppElement('#root');

moment.locale("en-GB");
const localizer = momentLocalizer(moment);

export interface BaseEvent {
    title: string,
    start: Date,
    end: Date,
    price: number,
}

export default function FullCalendar() {
    const [events, setEvents] = useState<BaseEvent[]>([]);
    const [modalIsOpen, setModalIsOpen] = useState(false);
    const [newEvent, setNewEvent] = useState({ start: null, end: null, title: '', price: 0 });

    const localizer = momentLocalizer(moment);

    const handleSelect = ({ start, end }) => {
        setNewEvent({ start, end, title: '' });
        setModalIsOpen(true);

    };

    const handleCloseModal = () => {
        setModalIsOpen(false);
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
            handleCloseModal();
        }
    };
    const handleChangeTitle = (e) => {
        setNewEvent({ ...newEvent, title: e.target.value });
    };

    const handleChangePrice = (e) => {
        setNewEvent({ ...newEvent, price: Number(e.target.value) });
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
            />
            <Modal
                isOpen={modalIsOpen}
                onRequestClose={handleCloseModal}
                contentLabel="New Event"
                styles={modalStyles}
            >
                <h2>New Event</h2>
                <Form onSubmit={(e) => {
                    e.preventDefault();
                    handleSave();
                }}>
                    <label>
                        Event Name:
                        <input
                            type="text"
                            value={newEvent.title}
                            onChange={handleChangeTitle}
                            required
                        />
                        <input
                            type="text"
                            value={newEvent.price}
                            onChange={handleChangePrice}
                            required
                        />
                    </label>
                    <br />
                    <button type="submit">Save</button>
                    <button type="button" onClick={handleCloseModal}>Cancel</button>
                </Form>
            </Modal>
        </div>
    );
}