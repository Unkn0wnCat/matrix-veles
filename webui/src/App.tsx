import React, {useState} from 'react';
import logo from './logo.svg';
import {Counter} from './features/counter/Counter';
import AxiosContext from "./context/axios";
import {AxiosError} from "axios";

function App() {
    const [test, setTest] = useState("Loading...")
    return (
        <div className="App">
            <AxiosContext.Consumer>
                {
                    axios => {

                        if(test == "Loading...") axios.get("/").then(res => {
                            setTest(JSON.stringify(res.data, null,  2))
                        }).catch((err: AxiosError) => {
                            setTest(JSON.stringify(err.response?.data))
                        })

                        return <pre>{test}</pre>
                    }
                }
            </AxiosContext.Consumer>
        </div>
    );
}

export default App;
