import React, {useState} from "react";
import {useNavigate, useOutlet} from "react-router-dom";

import styles from "./Entries.module.scss";
import {Trans, useTranslation} from "react-i18next";
import EntriesTable from "./EntriesTable";
import {PreloadedQuery, usePreloadedQuery} from "react-relay/hooks";
import {graphql} from "babel-plugin-relay/macro";
import {EntriesQuery} from "./__generated__/EntriesQuery.graphql";
import {X} from "lucide-react";
import {LoaderSuspense} from "../../../common/Loader";

type Props = {
    initialQueryRef: PreloadedQuery<EntriesQuery>,
}

const Entries = ({initialQueryRef}: Props) => {
    const outlet = useOutlet()
    const navigate = useNavigate()
    const {t} = useTranslation()

    const data = usePreloadedQuery(
            graphql`
                query EntriesQuery($first: String, $count: Int) {
                    ...EntriesTableFragment
                }
            `,
            initialQueryRef
    )

    const defaultTitle = t("entries.details", "Details")

    const [title, setTitle] = useState(defaultTitle)


    return <div className={styles.roomsContainer}>
        <div className={styles.roomsOverview + (outlet ? " "+styles.leaveSpace : "")}>
            <h1><Trans i18nKey={"entries.title"}>Entry Management</Trans></h1>

            <EntriesTable initialQueryRef={data}/>
        </div>

        <div className={styles.slideOver + (outlet ? " "+styles.active : "")}>
            <EntriesSlideOverTitleContext.Provider value={setTitle}>
                <div className={styles.slideOverHeader}>
                        <span>{title}</span>
                        <button onClick={() => navigate("/hashing/entries")}><X/></button>
                </div>
                <div className={styles.slideOverContent}>
                    <LoaderSuspense>
                        {outlet}
                    </LoaderSuspense>
                </div>
            </EntriesSlideOverTitleContext.Provider>
        </div>
    </div>
}

export const EntriesSlideOverTitleContext = React.createContext<React.Dispatch<React.SetStateAction<string>>|null>(null)

export default Entries