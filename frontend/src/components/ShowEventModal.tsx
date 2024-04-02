import {ModalStyled} from "./styles/FullCalendar.styled";
import React from "react";

interface ModalProps {
    modalIsOpen: boolean;
    handleCloseModal: () => void;
    eventId: string;
}

export default function ShowEventModal({modalIsOpen, handleCloseModal}: ModalProps) {
    return (
        <ModalStyled
            isOpen={modalIsOpen}
            onRequestClose={handleCloseModal}
            contentLabel="Event info"
        >
            <button type="button" className="btn btn-secondary" onClick={handleCloseModal}>Close</button>
        </ModalStyled>
    );
}
