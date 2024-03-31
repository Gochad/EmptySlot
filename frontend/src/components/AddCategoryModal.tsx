import React, {useState} from "react";
import { SketchPicker, ColorResult } from 'react-color';
import {Form, ModalStyled} from "./styles/FullCalendar.styled";
import {Category} from "../services/Categories";

interface ModalProps {
    modalIsOpen: boolean;
    item: Category;
    handleCloseModal: () => void;
    handleSave: () => void;
    handleChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
    handleColorChange: (category: Category) => void;
}
export default function AddCategoryModal({modalIsOpen, item, handleCloseModal, handleSave, handleChange, handleColorChange}: ModalProps) {
    const [color, setColor] = useState<string>('');

    const colorChange = (color: ColorResult) => {
        setColor(color.hex);
        handleColorChange({...item, color: color.hex})
    };

    return (
        <ModalStyled
            isOpen={modalIsOpen}
            onRequestClose={handleCloseModal}
            contentLabel="New Category"
        >
            <h2>New Event</h2>
            <Form onSubmit={(e) => {
                e.preventDefault();
                handleSave();
            }}>

                <div className="mb-3">
                    <label htmlFor="categoryName" className="form-label">Name:</label>
                    <input type="text" className="form-control" id="categoryName" name="name" value={item?.name || ''} onChange={handleChange} />
                </div>
                <SketchPicker color={color} onChangeComplete={colorChange} />

                <button type="submit" className="btn btn-primary">Save</button>
                <button type="button" className="btn btn-secondary" onClick={handleSave}>Cancel</button>
            </Form>
        </ModalStyled>
    );
}
