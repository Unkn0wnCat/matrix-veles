import React, {useState} from "react";

import styles from "./AuthViews.module.scss";

import {ReactComponent as Logo} from "../../logo.svg";
import {Link, useLocation} from "react-router-dom";
import {AuthLocationState} from "../../layouts/AuthLayout";
import {useAppDispatch} from "../../app/hooks";
import {Helmet} from "react-helmet";
import {Trans, useTranslation} from "react-i18next";
import {logIn} from "../../features/auth/authSlice";
import RegisterMutation from "../../mutations/RegisterMutation";
import {Key} from "lucide-react";

const RegisterView = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [matrix, setMatrix] = useState("");

    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);

    const location = useLocation()

    const locationState = location.state as AuthLocationState

    const dispatch = useAppDispatch()

    const {t} = useTranslation()

    const onSubmit = async () => {
        setLoading(true)
        setError("")

        try {
            const res = await RegisterMutation(username, password, matrix)

            const jwt = res.register;

            dispatch(logIn(jwt))
        } catch (e: any) {
            setError("An error occurred: " + e.source?.errors[0]?.message)
        } finally {
            setLoading(false)
        }
    }

    return <>
        <Logo width={64} height={64}/>
        <Helmet>
            <title>{t("auth:register.htmlTitle", "Register with Veles")}</title>
        </Helmet>

        <h1><Trans i18nKey={"auth:register.title"}>Register</Trans></h1>

        {loading && <div className={styles.loader}>
            <Key/>
            <span><Trans i18nKey={"auth:login.logging_in"}>Logging in...</Trans></span>
        </div>}

        {!loading && <>
            {error !== "" && <span className={styles.error}>{error}</span>}

            <form onSubmit={(e) => {
                e.preventDefault();
                onSubmit()
            }} className={styles.authForm}>
                <input onChange={(ev) => setUsername(ev.target.value)} value={username}
                       placeholder={t("auth:username", "Username")} autoCapitalize={"no"} autoCorrect={"no"}/>
                <input onChange={(ev) => setPassword(ev.target.value)} value={password}
                       placeholder={t("auth:password", "Password")} type={"password"} autoCapitalize={"no"}
                       autoCorrect={"no"}/>
                <input onChange={(ev) => setMatrix(ev.target.value)} value={matrix}
                       placeholder={t("auth:matrix_handle", "Matrix-Handle") + " (@user:matrix.org)"}
                       autoCapitalize={"no"} autoCorrect={"no"}/>
                <button onClick={() => onSubmit()}><Trans i18nKey={"auth:register.register"}>Register</Trans></button>
            </form>

            <Link to={"/auth/login"} className={styles.mindChangedLink} aria-label={t("auth:login.login", "Login")}
                  state={locationState}><Trans i18nKey={"auth:register.login_instead"}>I already have an account</Trans></Link>
        </>}
    </>
}

export default RegisterView;