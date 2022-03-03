import axios from "axios"

export const axiosDefault = axios.create({
    baseURL: "http://127.0.0.1:8123/api"
})