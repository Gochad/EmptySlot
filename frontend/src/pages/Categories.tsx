import React, {ChangeEvent, useEffect, useState} from 'react';
import Container from '@mui/material/Container';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import Navbar from "../components/Navbar";
import CategoriesService from "../services/Categories";
import {Category} from "../services/Categories";
import {errorPopup, successPopup} from "../services/utils";
import AddCategoryModal from "../components/AddCategoryModal";

export default function CategoriesScree() {
    const [categories, setCategories] = useState<Category[]>([]);
    const [newCategory, setNewCategory] = useState<Category>({ name: "", color: "" });
    const [modals, setModals] = useState({
        addCategoryModal: false,
    });

    const openModal = (modalName: string) => {
        setModals({ ...modals, [modalName]: true });
    };

    const closeModal = (modalName: string) => {
        setModals({ ...modals, [modalName]: false });
    };

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;

        setNewCategory(prevItem => ({
            ...prevItem,
            [name]: value
        }));
    };

    const handleSave = async () => {
        try {
            if (newCategory.name) {
                setCategories([
                    ...categories,
                    newCategory
                ]);
            }

            await CategoriesService.create(newCategory);
            successPopup(`category added`);
        } catch (error) {
            errorPopup(`error while saving new category: ${error}`);
        } finally {
            closeModal('addCategoryModal');
        }
    };

    const rerenderCategories = async () => {
        try {
            const items: Category[] = await CategoriesService.get();
            return items;
        } catch (error) {
            errorPopup(`problem with loading data: ${error}`);
        }
    };

    useEffect(() => {
        (async () => {
            try {
                const categories: Category[] | undefined = await rerenderCategories();
                if (categories) {
                    setCategories(categories);
                }
            } catch (error) {
                errorPopup(`error fetching events: ${error}`);
            }
        })();
    }, []);


    return (
        <>
            <Navbar />
            <Container maxWidth="md">
                <Typography variant="h4" component="h1" gutterBottom>
                    Service Categories
                </Typography>
                <Box mb={2}>
                    <Button variant="contained" color="primary" onClick={() => openModal('addCategoryModal')}>
                        Add category
                    </Button>
                </Box>
                <Grid container spacing={2}>
                    {categories.map((item) => (
                        <Grid item xs={12} sm={6} md={4} key={item.name}>
                            <Paper elevation={3} sx={{ padding: 2 , backgroundColor: item.color }}>
                                <Typography variant="h6" component="h2">
                                    {item.name}
                                </Typography>
                            </Paper>
                        </Grid>
                    ))}
                </Grid>
            </Container>
            <AddCategoryModal
                modalIsOpen={modals.addCategoryModal}
                handleCloseModal={() => closeModal("addCategoryModal")}
                handleSave={handleSave}
                item={newCategory}
                handleChange={handleChange}
                handleColorChange={setNewCategory}
            />
        </>
    );
}
