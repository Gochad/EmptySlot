import axios from "axios";
import {config} from "../config";

export interface Category {
    name: string,
    color: string,
}

export default class CategoriesService {
    static async get() {
        const response = await axios.get(`${config.API}${config.CATEGORIES}/`);
        const categories: Category[] = response.data;
        return categories;
    }

    static async create(data: Category) {
        await axios.post(`${config.API}${config.CATEGORIES}/`, data);
    }
}
