import React, {useEffect, useState} from "react";
import {Form, ModalStyled} from "./styles/FullCalendar.styled";
import { FormControl, MenuItem, Select, SelectChangeEvent } from '@mui/material';
import {BaseEvent} from "../services/events";
import CategoriesService, {Category} from "../services/Categories";
import {errorPopup} from "../services/utils";

interface ModalProps {
    modalIsOpen: boolean;
    handleCloseModal: () => void;
    handleSave: (categoryId: string) => void;
    handleChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
    event: BaseEvent;
}
export default function AddEventModal({modalIsOpen, handleCloseModal, handleSave, event, handleChange}: ModalProps) {
    const [categories, setCategories] = useState<Category[]>([]);

    const [selectedCategoryId, setselectedCategoryId] = React.useState('');

    const rerenderCategories = async () => {
        try {
            const items: Category[] = await CategoriesService.get();
            return items;
        } catch (error) {
            errorPopup(`problem with loading data: ${error}`);
        }
    };

    useEffect(() => {
        rerenderCategories().then(categories => {
            if (categories) {
                setCategories(categories);
            }
        }).catch(error => {
            errorPopup(`error fetching categories: ${error}`);
        });
    }, []);

    const handleChangeInFormControl = (event: SelectChangeEvent<string>) => {
        setselectedCategoryId(event.target.value);
    };

    return (
        <ModalStyled
            isOpen={modalIsOpen}
            onRequestClose={handleCloseModal}
            contentLabel="New Event"
        >
            <h2>New Event</h2>
            <Form onSubmit={(e) => {
                e.preventDefault();
                handleSave(selectedCategoryId);
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

                <FormControl fullWidth>
                    <label htmlFor="eventColor" className="form-label">Category:</label>
                    <Select
                        labelId="category-select-label"
                        id="category-select"
                        value={selectedCategoryId}
                        label="Category"
                        onChange={handleChangeInFormControl}
                    >
                        {categories.map((item) => (
                            <MenuItem key={item.id} value={item.id} sx={{ backgroundColor: item.color }}
                            >
                                {item.name}
                            </MenuItem>
                        ))}
                    </Select>
                </FormControl>

                <button type="submit" className="btn btn-primary">Save</button>
                <button type="button" className="btn btn-secondary" onClick={handleCloseModal}>Cancel</button>
            </Form>
        </ModalStyled>
    );
}
