import React from "react";

import styles from "./AuthLayout.module.scss";
import {Link, useLocation, useNavigate, useOutlet} from "react-router-dom";
import {UserPlus, User} from "lucide-react";

import {ReactComponent as Logo} from "../logo.svg";
import {useAppDispatch, useAppSelector} from "../app/hooks";
import {selectAuth} from "../features/auth/authSlice";

export type AuthLocationState = {
    location?: Location
}

const AuthLayout = () => {
    const outlet = useOutlet()
    const location = useLocation()

    const locationState = location.state as AuthLocationState

    const dispatch = useAppDispatch()
    const authState = useAppSelector(selectAuth)
    const navigate = useNavigate()


    if(authState.status == "logged_in") {
        if(locationState && locationState.location) {
            navigate(locationState.location, {replace: true})
        }
        navigate("/", {replace: true})
    }

    return <div className={styles.auth}>
        <div className={styles.background}/>
        <div className={styles.container}>
            <div className={styles.inner}>
                {outlet || <>
                    <Logo width={64} height={64} />
                    <h1>Matrix-Veles</h1>
                    <h2>Do we know each other?</h2>

                    <div className={styles.splitChoice}>
                        <Link to={"./login"} state={locationState}>
                            <User/>
                            <span>Yeah, let me log in</span>
                        </Link>
                        <Link to={"./register"} state={locationState}>
                            <UserPlus/>
                            <span>Nah, I'll sign up</span>
                        </Link>
                    </div>
                </>}
            </div>
            <footer>Veles WebUI | <a href={"https://veles.1in1.net/docs/intro"} target={"_blank"} rel={"noreferrer"}>Help</a> | <a href={"https://github.com/Unkn0wnCat/matrix-veles"} target={"_blank"} rel={"noreferrer"}>Source Code</a></footer>
        </div>
    </div>
}

export default AuthLayout