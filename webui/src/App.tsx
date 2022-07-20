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
import Dashboard from "./components/panel/dashboard/Dashboard";
import DashboardQueryGraphql, {DashboardQuery} from "./components/panel/dashboard/__generated__/DashboardQuery.graphql";
import Rooms from "./components/panel/rooms/Rooms";
import RoomsQueryGraphql, {RoomsQuery} from "./components/panel/rooms/__generated__/RoomsQuery.graphql";

function App() {
    const dispatch = useAppDispatch()
    const auth = useAppSelector(selectAuth)
    const environment = useRelayEnvironment();


    const [dashboardInitialState, loadQuery, disposeQuery] = useQueryLoader<DashboardQuery>(
            DashboardQueryGraphql
    )

    const [roomsInitialState, loadRoomsQuery, disposeRoomsQuery] = useQueryLoader<RoomsQuery>(
            RoomsQueryGraphql
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
            loadRoomsQuery({})
            return
        }

        disposeQuery()
        disposeRoomsQuery()
        environment.getStore().notify(undefined, true)
    }, [auth])


    return (
        <Routes>
            <Route path={"/auth"} element={<AuthLayout/>}>
                <Route path={"login"} element={<LoginView/>} />
                <Route path={"register"} element={<RegisterView/>} />
            </Route>
            <Route path={"/"} element={<PanelLayout/>}>
                <Route path={""} element={<RequireAuth>{dashboardInitialState && <Dashboard initialQueryRef={dashboardInitialState}/>}</RequireAuth>} />
                <Route path={"rooms"} element={<RequireAuth>{roomsInitialState && <Rooms initialQueryRef={roomsInitialState}/>}</RequireAuth>}>
                    <Route path={":id"} element={<h1>room detail</h1>} />
                </Route>
                <Route path={"hashing/lists"} element={<RequireAuth><h1>lists</h1></RequireAuth>}>
                    <Route path={":id"} element={<h1>list detail</h1>} />
                </Route>
                <Route path={"hashing/entries"} element={<RequireAuth><h1>entries</h1></RequireAuth>}>
                    <Route path={":id"} element={<h1>entry detail</h1>} />
                </Route>
            </Route>
        </Routes>
    );
}

export default App;
