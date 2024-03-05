import React, {useEffect, useState} from "react";
import { Calendar, momentLocalizer } from "react-big-calendar";
import moment from "moment";
import "react-big-calendar/lib/css/react-big-calendar.css";
import {Events} from "./events";
import {errorPopup, successPopup} from "./utils";
import Modal from 'react-modal';

moment.locale("en-GB");
const localizer = momentLocalizer(moment);

export interface BaseEvent {
    title: string,
    start: Date,
    end: Date
}

export default function FullCalendar() {
    const [events, setEvents] = useState<BaseEvent[]>([]);
    const [modalIsOpen, setModalIsOpen] = useState(false);
    const [newEvent, setNewEvent] = useState({ start: null, end: null, title: '' });

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
            successPopup(`Event added`);
        } catch (error) {
            errorPopup(`Error while confirmation: ${error}`);
        } finally {
            handleCloseModal();
        }
    };
    const handleChange = (e) => {
        setNewEvent({ ...newEvent, title: e.target.value });
    };

    const rerenderEvents = async () => {
        try {
            const events: BaseEvent[] = await Events.get();
            console.log(events);
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
            />
            <Modal
                isOpen={modalIsOpen}
                onRequestClose={handleCloseModal}
                contentLabel="New Event"
            >
                <h2>New Event</h2>
                <form onSubmit={(e) => {
                    e.preventDefault();
                    handleSave();
                }}>
                    <label>
                        Event Name:
                        <input
                            type="text"
                            value={newEvent.title}
                            onChange={handleChange}
                            required
                        />
                    </label>
                    <br />
                    <button type="submit">Save</button>
                    <button type="button" onClick={handleCloseModal}>Cancel</button>
                </form>
            </Modal>
        </div>
    );
}