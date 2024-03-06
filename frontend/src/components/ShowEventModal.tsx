import Modal from "react-modal";
import {Form, modalStyles} from "./styles/FullCalendar.styled";
import React from "react";


export default function ShowEventModal({modalIsOpen, handleCloseModal, title, price}) {
    return (
        <Modal
            isOpen={modalIsOpen}
            onRequestClose={handleCloseModal}
            contentLabel="Event info"
            styles={modalStyles}
        >
            <p>Title</p>
            <p>title</p>
            <p>`Price ${price}`</p>
            <button type="button" onClick={handleCloseModal}>Close</button>

        </Modal>
    );
}
