import React, {useState} from "react";

import styles from "./AuthViews.module.scss";

import {ReactComponent as Logo} from "../../logo.svg";
import {Link, useLocation} from "react-router-dom";
import {useAppDispatch} from "../../app/hooks";
import {logIn} from "../../features/auth/authSlice";

import {Key} from "lucide-react";
import {AuthLocationState} from "../../layouts/AuthLayout";
import {Trans, useTranslation} from "react-i18next";
import LoginMutation from "../../mutations/LoginMutation";
import Helmet from "react-helmet";

const LoginView = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);

    const dispatch = useAppDispatch()

    const location = useLocation();

    const locationState = location.state as AuthLocationState

    const {t} = useTranslation()

    const onSubmit = async () => {
        setLoading(true)
        setError("")

        try {
            const res = await LoginMutation(username, password)

            const jwt = res.login;

            dispatch(logIn(jwt))
        } catch (e: any) {
            setError("An error occurred: "+e.source?.errors[0]?.message)
        } finally {
            setLoading(false)
        }

        /*try {
            const res = await axiosDefault.post("/auth/login", {
                username,
                password
            })

            const jwt = res.data.token;

            dispatch(logIn(jwt))
        } catch (e: any) {
            if((e as AxiosError).isAxiosError) {
                const axErr = e as AxiosError

                setError(": "+axErr.response?.data.error)
            } else {
                setError("An unknown error occurred.")
            }
        } finally {
            setLoading(false)
        }*/
    }

    return <>
        <Logo width={64} height={64} />
        <Helmet>
            <title>{t("auth:login.htmlTitle", "Login to Veles")}</title>
        </Helmet>

        <h1><Trans i18nKey={"auth:login.title"}>Login</Trans></h1>

        { loading && <div className={styles.loader}>
            <Key/>
            <span><Trans i18nKey={"auth:login.logging_in"}>Logging in...</Trans></span>
        </div>}

        { !loading && <>
            {error !== "" && <span className={styles.error}>{error}</span>}

            <form onSubmit={(e) => {e.preventDefault(); onSubmit()}} className={styles.authForm}>
                <input onChange={(ev) => setUsername(ev.target.value)} value={username} placeholder={t("auth:username", "Username")} />
                <input onChange={(ev) => setPassword(ev.target.value)} value={password} placeholder={t("auth:password", "Password")} type={"password"} />
                <button onClick={() => onSubmit()}><Trans i18nKey={"auth:login.login"}>Login</Trans></button>
            </form>

            <Link to={"/auth/register"} className={styles.mindChangedLink} aria-label={t("auth:register.register", "Register")} state={locationState}><Trans i18nKey={"auth:login.register_instead"}>I don't have an account</Trans></Link>
        </>}
    </>
}

export default LoginView;