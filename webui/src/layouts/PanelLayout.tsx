import React from "react";

import {Link, NavLink, useOutlet} from "react-router-dom";
import {Home, List, ClipboardList, ExternalLink} from "lucide-react";

import {ReactComponent as Logo} from "../logo.svg";

import styles from "./PanelLayout.module.scss";

const PanelLayout = () => {
    const outlet = useOutlet();

    return <div className={styles.panel}>
        <a href={"#main"} className={styles.skipToContent}>Jump to Content</a>
        <a href={"#navigation"} className={styles.skipToContent}>Jump to Navigation</a>
        <div className={styles.topBar}>
            <Link to={"/"} className={styles.logo}><Logo/> <span>Matrix-Veles</span></Link>
            <a href={"https://veles.1in1.net/docs/intro"} target={"_blank"} rel={"noreferrer"}><ExternalLink/> <span>Documentation</span></a>
        </div>
        <div className={styles.content}>
            <nav id={"navigation"}>
                <NavLink to={"/"}><Home/><span>Dashboard</span></NavLink>
                <NavLink to={"/hashing/lists"}><List/><span>Lists</span></NavLink>
                <NavLink to={"/hashing/entries"}><ClipboardList/><span>Entries</span></NavLink>
            </nav>
            <main id={"main"}>
                {outlet}
            </main>
        </div>
    </div>
}

export default PanelLayout