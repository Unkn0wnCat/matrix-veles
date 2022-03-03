import type { RootState } from "./store";

export const loadState = () => {
    try {
        const serializedState = localStorage.getItem("velesAppState");

        if (serializedState === null) {
            return undefined;
        }

        return JSON.parse(serializedState);
    } catch (err) {
        return undefined;
    }
};

export const saveState = (state: RootState) => {
    try {
        const serializedState = JSON.stringify({
            auth: state.auth
        });
        localStorage.setItem("velesAppState", serializedState);
    } catch {}
};