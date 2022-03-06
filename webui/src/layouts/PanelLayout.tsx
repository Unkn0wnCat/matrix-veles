import React, {useState} from "react";

import {Link, NavLink, useOutlet} from "react-router-dom";
import {Home, List, ClipboardList, ExternalLink, ChevronRight, MessageSquare} from "lucide-react";

import {ReactComponent as Logo} from "../logo.svg";

import styles from "./PanelLayout.module.scss";
import {Trans} from "react-i18next";

const PanelLayout = () => {
    const outlet = useOutlet();
    const [hashingExpanded, setHashingExpanded] = useState(false)

    return <div className={styles.panel}>
        <a href={"#main"} className={styles.skipToContent}><Trans i18nKey={"panel:jump_to_content"}>Jump to Content</Trans></a>
        <a href={"#navigation"} className={styles.skipToContent}><Trans i18nKey={"panel:jump_to_navigation"}>Jump to Navigation</Trans></a>
        <div className={styles.topBar}>
            <Link to={"/"} className={styles.logo}><Logo/> <span>Matrix-Veles</span></Link>
            <a href={"https://veles.1in1.net/docs/intro"} target={"_blank"} rel={"noreferrer"}><ExternalLink/> <span><Trans i18nKey={"panel:documentation"}>Documentation</Trans></span></a>
        </div>
        <div className={styles.content}>
            <nav id={"navigation"}>
                <NavLink to={"/"}><Home/><span><Trans i18nKey={"panel:navigation.dashboard"}>Dashboard</Trans></span></NavLink>
                <NavLink to={"/rooms"}><MessageSquare/><span><Trans i18nKey={"panel:navigation.rooms"}>My Rooms</Trans></span></NavLink>
                <div className={styles.dropdown + (hashingExpanded?" "+styles.expanded:"")}>
                    <button onClick={() => setHashingExpanded(!hashingExpanded)}>
                        <ChevronRight/> <span><Trans i18nKey={"panel:navigation.hash_checker"}>Hash-Checker</Trans></span>
                    </button>
                    <NavLink to={"/hashing/lists"}><List/><span><Trans i18nKey={"panel:navigation.hash_lists"}>Lists</Trans></span></NavLink>
                    <NavLink to={"/hashing/entries"}><ClipboardList/><span><Trans i18nKey={"panel:navigation.hash_entries"}>Entries</Trans></span></NavLink>
                </div>
            </nav>
            <main id={"main"}>
                {outlet}
            </main>
        </div>
    </div>
}

export default PanelLayout