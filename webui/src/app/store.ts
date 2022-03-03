import {configureStore, ThunkAction, Action} from '@reduxjs/toolkit';
import counterReducer from '../features/counter/counterSlice';
import authReducer from "../features/auth/authSlice";
import {loadState, saveState} from "./localStorage";

const preloadedState = loadState()

export const store = configureStore({
    reducer: {
        counter: counterReducer,
        auth: authReducer,
    },
    preloadedState
});

store.subscribe(() => {
    saveState(store.getState())
})

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<ReturnType,
    RootState,
    unknown,
    Action<string>>;
