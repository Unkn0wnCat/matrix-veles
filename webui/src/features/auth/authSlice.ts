import {createAsyncThunk, createSlice, PayloadAction} from '@reduxjs/toolkit';
import {RootState, AppThunk} from '../../app/store';
import broadcastChannel, { BroadcastMessage, sendMessage } from "../../app/broadcastChannel";

export interface AuthState {
    jwt: string|null;
    status: 'logged_in' | 'logged_out' | 'awaiting_verification';
}

const initialState: AuthState = {
    jwt: null,
    status: 'logged_out',
};

type BroadcastPayload = {
    status: "logged_in" | "logged_out" | "awaiting_verification",
    jwt: string|null
}

export const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
        logIn: (state, action: PayloadAction<string>) => {
            state.jwt = action.payload
            state.status = "logged_in"

            sendMessage({
                action: "updateAuth",
                payload: {
                    status: state.status,
                    jwt: state.jwt
                }
            })
        },
        logOut: (state) => {
            state.jwt = null
            state.status = "logged_out"

            sendMessage({
                action: "updateAuth",
                payload: {
                    status: state.status,
                    jwt: state.jwt
                }
            })
        },
        awaitVerification: (state, action: PayloadAction<string>) => {
            state.jwt = action.payload
            state.status = "awaiting_verification"

            sendMessage({
                action: "updateAuth",
                payload: {
                    status: state.status,
                    jwt: state.jwt
                }
            })
        },
        receiveAuthUpdate: (state, action: PayloadAction<BroadcastMessage<BroadcastPayload>>) => {
            if(action.payload.action !== "updateAuth") return;

            state.jwt = action.payload.payload.jwt
            state.status = action.payload.payload.status
        },
    },
});


export const {logIn, logOut, awaitVerification, receiveAuthUpdate} = authSlice.actions;

export const selectAuth = (state: RootState) => state.auth;

export default authSlice.reducer;
