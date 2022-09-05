import React from "react";
import {Loader as LoaderIcon} from "lucide-react";

import styles from "./Loader.module.scss";
import {Trans} from "react-i18next";

const Loader = () => {
    return <div className={styles.loader}>
        <LoaderIcon/>
        <span><Trans i18nKey={"common.loading"}>Loading...</Trans></span>
    </div>
}

export const LoaderSuspense = (props: React.PropsWithChildren<{}>) => <React.Suspense fallback={<Loader/>} children={props.children} />

//export const LoaderSuspense = (props: React.PropsWithChildren<{}>) => <Loader/>

export default Loader;