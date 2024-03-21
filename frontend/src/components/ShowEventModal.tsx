import Modal from "react-modal";
import {modalStyles} from "./styles/FullCalendar.styled";
import React from "react";
import {errorPopup} from "./utils";
import {Reservation} from "./reservation";


export default function ShowEventModal({modalIsOpen, handleCloseModal, eventId}) {
    const createPaymentLink = async () => {
        try {
            const link = await Reservation.pay(eventId);
        } catch (error) {
            errorPopup(`problem with creating payment link: ${error}`);
        }
    };

    return (
        <Modal
            isOpen={modalIsOpen}
            onRequestClose={handleCloseModal}
            contentLabel="Event info"
            styles={modalStyles}
        >
            <button type="button" onClick={createPaymentLink}>Pay</button>
            <button type="button" onClick={handleCloseModal}>Close</button>

        </Modal>
    );
}
