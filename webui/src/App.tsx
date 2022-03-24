import React from 'react';
import { Routes, Route } from "react-router-dom";
import AuthLayout from "./layouts/AuthLayout";
import LoginView from "./components/auth/LoginView";
import RegisterView from "./components/auth/RegisterView";
import RequireAuth from "./features/auth/RequireAuth";
import {useAppDispatch} from "./app/hooks";
import broadcastChannel from "./app/broadcastChannel";
import {logOut, receiveAuthUpdate} from "./features/auth/authSlice";
import PanelLayout from "./layouts/PanelLayout";
import {Trans, useTranslation} from "react-i18next";

function App() {
    const dispatch = useAppDispatch()

    // This needs to be here to prevent a weird bug
    useTranslation()

    broadcastChannel.on("message", (ev) => {
        if(ev.action === "updateAuth") {
            dispatch(receiveAuthUpdate(ev))
        }
    })

    return (
        <Routes>
            <Route path={"/auth"} element={<AuthLayout/>}>
                <Route path={"login"} element={<LoginView/>} />
                <Route path={"register"} element={<RegisterView/>} />
            </Route>
            <Route path={"/"} element={<PanelLayout/>}>
                <Route path={""} element={<RequireAuth><h1><Trans i18nKey={"test"}>Test</Trans></h1> <button onClick={() => {
                    dispatch(logOut())
                }
                }>Log out</button></RequireAuth>} />
                <Route path={"rooms"} element={<RequireAuth><h1>rooms</h1></RequireAuth>} />
                <Route path={"hashing/lists"} element={<RequireAuth><h1>lists</h1></RequireAuth>} />
                <Route path={"hashing/entries"} element={<RequireAuth><h1>entries</h1></RequireAuth>} />
            </Route>
        </Routes>
    );
}

export default App;
