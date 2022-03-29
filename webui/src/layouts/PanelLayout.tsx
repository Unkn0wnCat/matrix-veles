import React, {useState} from "react";

import {Link, NavLink, useOutlet} from "react-router-dom";
import {Home, List, ClipboardList, BookOpen, ChevronRight, MessageSquare, LogOut} from "lucide-react";
import { Squash as Hamburger } from 'hamburger-react';

import {ReactComponent as Logo} from "../logo.svg";

import styles from "./PanelLayout.module.scss";
import {Trans, useTranslation} from "react-i18next";
import {useAppDispatch} from "../app/hooks";
import {logOut} from "../features/auth/authSlice";

const PanelLayout = () => {
    const {t} = useTranslation()
    const outlet = useOutlet();
    const [hashingExpanded, setHashingExpanded] = useState(false)
    const [sideNavExpanded, setSideNavExpanded] = useState(true)

    const dispatch = useAppDispatch()

    return <div className={styles.panel}>
        <a href={"#main"} className={styles.skipToContent}><Trans i18nKey={"panel:jump_to_content"}>Jump to Content</Trans></a>
        <a href={"#navigation"} className={styles.skipToContent}><Trans i18nKey={"panel:jump_to_navigation"}>Jump to Navigation</Trans></a>
        <div className={styles.topBar}>
            <div className={styles.hamburger}>
                <Hamburger toggle={setSideNavExpanded} toggled={sideNavExpanded} rounded={true} duration={.25} distance={"lg"} />
            </div>
            <Link to={"/"} className={styles.logo}><Logo/> <span>Matrix-Veles</span></Link>
            <a href={"https://veles.1in1.net/docs/intro"} target={"_blank"} rel={"noreferrer"} title={t("panel:documentation", "Documentation")}><BookOpen/><span> <Trans i18nKey={"panel:documentation"}>Documentation</Trans></span></a>
            <button onClick={() => {dispatch(logOut())}} title={t("panel:logout", "Logout")}><LogOut/><span> <Trans i18nKey={"panel:logout"}>Logout</Trans></span></button>
        </div>
        <div className={styles.content}>
            <div className={styles.navUnderlay + (sideNavExpanded ? " "+styles.show : "")} onClick={() => setSideNavExpanded(false)} />
            <nav id={"navigation"} className={sideNavExpanded ? styles.expanded : ""}>
                <NavLink to={"/"} onClick={() => setSideNavExpanded(false)}><Home/><span><Trans i18nKey={"panel:navigation.dashboard"}>Dashboard</Trans></span></NavLink>
                <NavLink to={"/rooms"} onClick={() => setSideNavExpanded(false)}><MessageSquare/><span><Trans i18nKey={"panel:navigation.rooms"}>My Rooms</Trans></span></NavLink>
                <div className={styles.dropdown + (hashingExpanded?" "+styles.expanded:"")}>
                    <button onClick={() => setHashingExpanded(!hashingExpanded)}>
                        <ChevronRight/> <span><Trans i18nKey={"panel:navigation.hash_checker"}>Hash-Checker</Trans></span>
                    </button>
                    <NavLink to={"/hashing/lists"} onClick={() => setSideNavExpanded(false)}><List/><span><Trans i18nKey={"panel:navigation.hash_lists"}>Lists</Trans></span></NavLink>
                    <NavLink to={"/hashing/entries"} onClick={() => setSideNavExpanded(false)}><ClipboardList/><span><Trans i18nKey={"panel:navigation.hash_entries"}>Entries</Trans></span></NavLink>
                </div>
            </nav>
            <main id={"main"}>
                <React.Suspense fallback={<span>Fetching data...</span>}>
                    {outlet}
                </React.Suspense>
            </main>
        </div>
    </div>
}

export default PanelLayout