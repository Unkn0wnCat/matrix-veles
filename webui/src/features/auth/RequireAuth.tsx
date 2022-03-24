import React from "react";
import {useAppSelector} from "../../app/hooks";
import {selectAuth} from "./authSlice";
import {useLocation} from "react-router-dom";
import {Navigate} from "react-router-dom";

const RequireAuth = ({children}: React.PropsWithChildren<{}>): JSX.Element => {
    const authState = useAppSelector(selectAuth)
    const location = useLocation()

    if(authState.status !== "logged_in") {
        return <Navigate to="/auth" state={{from: location}} replace />
    }

    return <>{children}</>
}

export default RequireAuth