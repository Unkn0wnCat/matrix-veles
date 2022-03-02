import axios from "axios"
import React from "react";

export const axiosDefault = axios.create({
    baseURL: "http://127.0.0.1:8123/api"
})

const AxiosContext = React.createContext(axiosDefault);

export default AxiosContext