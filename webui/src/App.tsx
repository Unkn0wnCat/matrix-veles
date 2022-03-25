import React, {useEffect} from 'react';
import { Routes, Route } from "react-router-dom";
import AuthLayout from "./layouts/AuthLayout";
import LoginView from "./components/auth/LoginView";
import RegisterView from "./components/auth/RegisterView";
import RequireAuth from "./features/auth/RequireAuth";
import {useAppDispatch, useAppSelector} from "./app/hooks";
import broadcastChannel from "./app/broadcastChannel";
import {logOut, receiveAuthUpdate, selectAuth} from "./features/auth/authSlice";
import PanelLayout from "./layouts/PanelLayout";
import {useTranslation} from "react-i18next";
import {
    useQueryLoader, useRelayEnvironment,
} from 'react-relay/hooks';
import Dashboard from "./components/dashboard/Dashboard";
import DashboardQueryGraphql, {DashboardQuery} from "./components/dashboard/__generated__/DashboardQuery.graphql";

function App() {
    const dispatch = useAppDispatch()
    const auth = useAppSelector(selectAuth)
    const environment = useRelayEnvironment();


    const [dashboardInitialState, loadQuery, disposeQuery] = useQueryLoader<DashboardQuery>(
            DashboardQueryGraphql
    )

    // This needs to be here to prevent a weird bug
    useTranslation()

    broadcastChannel.on("message", (ev) => {
        if(ev.action === "updateAuth") {
            dispatch(receiveAuthUpdate(ev))
        }
    })

    useEffect(() => {
        if(auth.jwt !== null) {
            loadQuery({})
            return
        }

        disposeQuery()
        environment.getStore().notify(undefined, true)
    }, [auth])


    return (
        <Routes>
            <Route path={"/auth"} element={<AuthLayout/>}>
                <Route path={"login"} element={<LoginView/>} />
                <Route path={"register"} element={<RegisterView/>} />
            </Route>
            <Route path={"/"} element={<PanelLayout/>}>
                <Route path={""} element={<RequireAuth>{/*<h1><Trans i18nKey={"test"}>Test</Trans></h1> <button onClick={() => {
                    dispatch(logOut())
                }
                }>Log out</button> <p>{
                    JSON.stringify(data.self)
                }</p>*/}{dashboardInitialState && <Dashboard initialQueryRef={dashboardInitialState}/>}</RequireAuth>} />
                <Route path={"rooms"} element={<RequireAuth><h1>rooms</h1></RequireAuth>} />
                <Route path={"hashing/lists"} element={<RequireAuth><h1>lists</h1></RequireAuth>} />
                <Route path={"hashing/entries"} element={<RequireAuth><h1>entries</h1></RequireAuth>} />
            </Route>
        </Routes>
    );
}

export default App;
