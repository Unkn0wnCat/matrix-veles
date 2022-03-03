import React, {useRef, useState} from "react";

import styles from "./AuthViews.module.scss";

import {ReactComponent as Logo} from "../../logo.svg";
import {Link, useLocation, useNavigate} from "react-router-dom";
import {AuthLocationState} from "../../layouts/AuthLayout";
import {useAppDispatch, useAppSelector} from "../../app/hooks";
import {selectAuth} from "../../features/auth/authSlice";

const RegisterView = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [matrix, setMatrix] = useState("");

    const location = useLocation()

    const locationState = location.state as AuthLocationState

    const dispatch = useAppDispatch()

    const onSubmit = () => {
        console.log(username, password)
    }

    return <>
        <Logo width={64} height={64} />
        <h1>Register</h1>

        <form onSubmit={(e) => {e.preventDefault(); onSubmit()}} className={styles.authForm}>
            <input onChange={(ev) => setUsername(ev.target.value)} value={username} placeholder={"Username"} autoCapitalize={"no"} autoCorrect={"no"} />
            <input onChange={(ev) => setPassword(ev.target.value)} value={password} placeholder={"Password"} type={"password"} autoCapitalize={"no"} autoCorrect={"no"} />
            <input onChange={(ev) => setMatrix(ev.target.value)} value={matrix} placeholder={"Matrix-Handle (@user:matrix.org)"} autoCapitalize={"no"} autoCorrect={"no"} />
            <button onClick={() => onSubmit()}>Register</button>
        </form>

        <Link to={"/auth/login"} className={styles.mindChangedLink} aria-label={"Register"} state={locationState}>I already have an account</Link>
    </>
}

export default RegisterView;