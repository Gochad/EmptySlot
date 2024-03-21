import React from "react";
import Modal from "react-modal";
import {Form, modalStyles} from "./styles/FullCalendar.styled";

export default function AddEventModal({modalIsOpen, handleCloseModal, handleSave, event, handleChange}) {
    return (
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
                    <input name="title" value={event?.title || ''} onChange={handleChange} />
                    <input name="price" type="number" value={event?.price || ''} onChange={handleChange} />
                    <input name="description" value={event?.description || ''} onChange={handleChange} />
                </label>
                <br />
                <button type="submit">Save</button>
                <button type="button" onClick={handleCloseModal}>Cancel</button>
            </Form>
        </Modal>
    );
}
