import React from "react";
import Modal from "react-modal";
import {Form, modalStyles} from "./styles/FullCalendar.styled";

export default function AddEventModal({modalIsOpen, handleCloseModal, handleSave, title, price, handleChangeTitle, handleChangePrice}) {
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
                    <input
                        type="text"
                        value={title}
                        onChange={handleChangeTitle}
                        required
                    />
                    <input
                        type="text"
                        value={price}
                        onChange={handleChangePrice}
                        required
                    />
                </label>
                <br />
                <button type="submit">Save</button>
                <button type="button" onClick={handleCloseModal}>Cancel</button>
            </Form>
        </Modal>
    );
}
