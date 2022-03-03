import React, {useRef, useState} from "react";

import styles from "./AuthViews.module.scss";

import {ReactComponent as Logo} from "../../logo.svg";
import {Link, useLocation, useNavigate} from "react-router-dom";
import {axiosDefault} from "../../context/axios";
import {AxiosError} from "axios";
import {useAppDispatch, useAppSelector} from "../../app/hooks";
import {logIn, selectAuth} from "../../features/auth/authSlice";

import {Key} from "lucide-react";
import {AuthLocationState} from "../../layouts/AuthLayout";

const LoginView = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);

    const dispatch = useAppDispatch()

    const location = useLocation();

    const locationState = location.state as AuthLocationState

    const onSubmit = async () => {
        setLoading(true)
        setError("")

        try {
            const res = await axiosDefault.post("/auth/login", {
                username,
                password
            })

            const jwt = res.data.token;

            dispatch(logIn(jwt))
        } catch (e: any) {
            if((e as AxiosError).isAxiosError) {
                const axErr = e as AxiosError

                setError("Server returned error:    "+axErr.response?.data.error)
            } else {
                setError("An unknown error occurred.")
            }
        } finally {
            setLoading(false)
        }
    }

    return <>
        <Logo width={64} height={64} />
        <h1>Login</h1>

        { loading && <div className={styles.loader}>
            <Key/>
            <span>Logging in...</span>
        </div>}

        { !loading && <>
            {error !== "" && <span className={styles.error}>{error}</span>}

            <form onSubmit={(e) => {e.preventDefault(); onSubmit()}} className={styles.authForm}>
                <input onChange={(ev) => setUsername(ev.target.value)} value={username} placeholder={"Username"} />
                <input onChange={(ev) => setPassword(ev.target.value)} value={password} placeholder={"Password"} type={"password"} />
                <button onClick={() => onSubmit()}>Login</button>
            </form>

            <Link to={"/auth/register"} className={styles.mindChangedLink} aria-label={"Register"} state={locationState}>I don't have an account</Link>
        </>}
    </>
}

export default LoginView;