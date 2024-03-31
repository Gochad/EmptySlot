import React from "react";
import {Form, ModalStyled} from "./styles/FullCalendar.styled";
import {BaseEvent} from "../services/events";

interface ModalProps {
    modalIsOpen: boolean;
    handleCloseModal: () => void;
    handleSave: () => void;
    handleChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
    event: BaseEvent;
}
export default function AddEventModal({modalIsOpen, handleCloseModal, handleSave, event, handleChange}: ModalProps) {
    return (
        <ModalStyled
            isOpen={modalIsOpen}
            onRequestClose={handleCloseModal}
            contentLabel="New Event"
        >
            <h2>New Event</h2>
            <Form onSubmit={(e) => {
                e.preventDefault();
                handleSave();
            }}>

                <div className="mb-3">
                    <label htmlFor="eventName" className="form-label">Event Name:</label>
                    <input type="text" className="form-control" id="eventName" name="title" value={event?.title || ''} onChange={handleChange} />
                </div>

                <div className="mb-3">
                    <label htmlFor="eventPrice" className="form-label">Price:</label>
                    <input type="number" className="form-control" id="eventPrice" name="price" value={event?.price || ''} onChange={handleChange} />
                </div>

                <div className="mb-3">
                    <label htmlFor="eventDescription" className="form-label">Description:</label>
                    <input type="text" className="form-control" id="eventDescription" name="description" value={event?.description || ''} onChange={handleChange} />
                </div>

                <button type="submit" className="btn btn-primary">Save</button>
                <button type="button" className="btn btn-secondary" onClick={handleCloseModal}>Cancel</button>
            </Form>
        </ModalStyled>
    );
}
