import React, {useEffect} from 'react';
import { Routes, Route } from "react-router-dom";
import AuthLayout from "./layouts/AuthLayout";
import LoginView from "./components/auth/LoginView";
import RegisterView from "./components/auth/RegisterView";
import RequireAuth from "./features/auth/RequireAuth";
import {useAppDispatch, useAppSelector} from "./app/hooks";
import broadcastChannel from "./app/broadcastChannel";
import {receiveAuthUpdate, selectAuth} from "./features/auth/authSlice";
import PanelLayout from "./layouts/PanelLayout";
import {useTranslation} from "react-i18next";
import {
    useQueryLoader, useRelayEnvironment,
} from 'react-relay/hooks';
import Dashboard from "./components/panel/dashboard/Dashboard";
import DashboardQueryGraphql, {DashboardQuery} from "./components/panel/dashboard/__generated__/DashboardQuery.graphql";
import Rooms from "./components/panel/rooms/Rooms";
import RoomsQueryGraphql, {RoomsQuery} from "./components/panel/rooms/__generated__/RoomsQuery.graphql";
import RoomDetailQueryGraphql, {RoomDetailQuery} from "./components/panel/rooms/__generated__/RoomDetailQuery.graphql";
import RoomDetail from "./components/panel/rooms/RoomDetail";
import ListsQueryGraphql, {ListsQuery} from "./components/panel/hashing/lists/__generated__/ListsQuery.graphql";
import ListDetailQueryGraphql, {ListDetailQuery} from "./components/panel/hashing/lists/__generated__/ListDetailQuery.graphql";
import Lists from "./components/panel/hashing/lists/Lists";
import ListDetail from "./components/panel/hashing/lists/ListDetail";
import EntriesQueryGraphql, {EntriesQuery} from "./components/panel/hashing/entries/__generated__/EntriesQuery.graphql";
import EntryDetailQueryGraphql, {EntryDetailQuery} from "./components/panel/hashing/entries/__generated__/EntryDetailQuery.graphql";
import Entries from "./components/panel/hashing/entries/Entries";
import EntryDetail from "./components/panel/hashing/entries/EntryDetail";

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

    const [roomDetailInitialState, loadRoomDetailQuery, disposeRoomDetailQuery] = useQueryLoader<RoomDetailQuery>(
            RoomDetailQueryGraphql
    )

    const [listsInitialState, loadListsQuery, disposeListsQuery] = useQueryLoader<ListsQuery>(
            ListsQueryGraphql
    )

    const [listDetailInitialState, loadListDetailQuery, disposeListDetailQuery] = useQueryLoader<ListDetailQuery>(
            ListDetailQueryGraphql
    )

    const [entriesInitialState, loadEntriesQuery, disposeEntriesQuery] = useQueryLoader<EntriesQuery>(
            EntriesQueryGraphql
    )

    const [entryDetailInitialState, loadEntryDetailQuery, disposeEntryDetailQuery] = useQueryLoader<EntryDetailQuery>(
            EntryDetailQueryGraphql
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
            loadListsQuery({})
            loadEntriesQuery({})
            return
        }

        disposeQuery()
        disposeRoomsQuery()
        disposeListsQuery()
        disposeEntriesQuery()
        environment.getStore().notify(undefined, true)
    }, [auth, disposeQuery, disposeRoomsQuery, environment, loadQuery, loadRoomsQuery, loadListsQuery, disposeListsQuery, loadEntriesQuery, disposeEntriesQuery])


    return (
        <Routes>
            <Route path={"/auth"} element={<AuthLayout/>}>
                <Route path={"login"} element={<LoginView/>} />
                <Route path={"register"} element={<RegisterView/>} />
            </Route>
            <Route path={"/"} element={<PanelLayout/>}>
                <Route path={""} element={<RequireAuth>{dashboardInitialState && <Dashboard initialQueryRef={dashboardInitialState}/>}</RequireAuth>} />
                <Route path={"rooms"} element={<RequireAuth>{roomsInitialState && <Rooms initialQueryRef={roomsInitialState}/>}</RequireAuth>}>
                    <Route path={":id"} element={<RequireAuth><RoomDetail initialQueryRef={roomDetailInitialState} fetch={loadRoomDetailQuery} dispose={disposeRoomDetailQuery}/></RequireAuth>} />
                </Route>
                <Route path={"hashing/lists"} element={<RequireAuth>{listsInitialState && <Lists initialQueryRef={listsInitialState}/>}</RequireAuth>}>
                    <Route path={":id"} element={<RequireAuth><ListDetail initialQueryRef={listDetailInitialState} fetch={loadListDetailQuery} dispose={disposeListDetailQuery}/></RequireAuth>} />
                </Route>
                <Route path={"hashing/entries"} element={<RequireAuth>{entriesInitialState && <Entries initialQueryRef={entriesInitialState}/>}</RequireAuth>}>
                    <Route path={":id"} element={<RequireAuth><EntryDetail initialQueryRef={entryDetailInitialState} fetch={loadEntryDetailQuery} dispose={disposeEntryDetailQuery}/></RequireAuth>} />
                </Route>
            </Route>
        </Routes>
    );
}

export default App;
