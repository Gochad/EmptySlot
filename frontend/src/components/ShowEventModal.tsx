import {ModalStyled} from "./styles/FullCalendar.styled";
import React from "react";
import {errorPopup} from "./utils";
import {Reservation} from "./reservation";

interface ModalProps {
    modalIsOpen: boolean;
    handleCloseModal: () => void;
    eventId: string;
}

export default function ShowEventModal({modalIsOpen, handleCloseModal, eventId}: ModalProps) {
    const createPaymentLink = async () => {
        try {
            const link = await Reservation.pay(eventId);
            window.location.href = link;
        } catch (error) {
            errorPopup(`problem with creating payment link: ${error}`);
        }
    };

    return (
        <ModalStyled
            isOpen={modalIsOpen}
            onRequestClose={handleCloseModal}
            contentLabel="Event info"
        >
            <button type="submit" className="btn btn-primary" onClick={createPaymentLink}>Pay</button>
            <button type="button" className="btn btn-secondary" onClick={handleCloseModal}>Close</button>
        </ModalStyled>
    );
}
