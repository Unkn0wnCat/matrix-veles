import React, {useEffect} from "react";

import styles from "./AuthLayout.module.scss";
import {Link, useLocation, useNavigate, useOutlet} from "react-router-dom";
import {UserPlus, User} from "lucide-react";

import {ReactComponent as Logo} from "../logo.svg";
import {useAppSelector} from "../app/hooks";
import {selectAuth} from "../features/auth/authSlice";
import {Trans} from "react-i18next";

export type AuthLocationState = {
    location?: Location
}

const AuthLayout = () => {
    const outlet = useOutlet()
    const location = useLocation()

    const locationState = location.state as AuthLocationState

    const authState = useAppSelector(selectAuth)
    const navigate = useNavigate()


    useEffect(() => {
        if(authState.status === "logged_in") {
            if(locationState && locationState.location) {
                navigate(locationState.location, {replace: true})
            }
            navigate("/", {replace: true})
        }
    }, [authState, locationState, navigate])

    return <div className={styles.auth}>
        <div className={styles.background}/>
        <div className={styles.container}>
            <div className={styles.inner}>
                {outlet || <>
                    <Logo width={64} height={64} />
                    <h1>Matrix-Veles</h1>
                    <h2><Trans i18nKey={"auth:selector.question"}>Do we know each other?</Trans></h2>

                    <div className={styles.splitChoice}>
                        <Link to={"./login"} state={locationState}>
                            <User/>
                            <span><Trans i18nKey={"auth:selector.login"}>Yeah, let me log in</Trans></span>
                        </Link>
                        <Link to={"./register"} state={locationState}>
                            <UserPlus/>
                            <span><Trans i18nKey={"auth:selector.register"}>Nah, I'll sign up</Trans></span>
                        </Link>
                    </div>
                </>}
            </div>
            <footer>Veles WebUI | <a href={"https://veles.1in1.net/docs/intro"} target={"_blank"} rel={"noreferrer"}><Trans i18nKey={"auth:help"}>Help</Trans></a> | <a href={"https://github.com/Unkn0wnCat/matrix-veles"} target={"_blank"} rel={"noreferrer"}><Trans i18nKey={"auth:source"}>Source Code</Trans></a></footer>
        </div>
    </div>
}

export default AuthLayout