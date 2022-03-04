import React, {useState} from 'react';
import { Routes, Route, Link } from "react-router-dom";
import AuthLayout from "./layouts/AuthLayout";
import LoginView from "./components/auth/LoginView";
import RegisterView from "./components/auth/RegisterView";
import RequireAuth from "./features/auth/RequireAuth";
import {useAppDispatch} from "./app/hooks";
import broadcastChannel from "./app/broadcastChannel";
import {logOut, receiveAuthUpdate} from "./features/auth/authSlice";

function App() {
    const dispatch = useAppDispatch()

    broadcastChannel.on("message", (ev) => {
        if(ev.action == "updateAuth") {
            dispatch(receiveAuthUpdate(ev))
        }
    })

    return (
        <Routes>
            <Route path={"/auth"} element={<AuthLayout/>}>
                <Route path={"login"} element={<LoginView/>} />
                <Route path={"register"} element={<RegisterView/>} />
            </Route>
            <Route path={"/"} element={<RequireAuth><h1>hi</h1> <button onClick={() => {
                dispatch(logOut())
            }
            }>Log out</button></RequireAuth>}/>
        </Routes>
    );
}

export default App;
