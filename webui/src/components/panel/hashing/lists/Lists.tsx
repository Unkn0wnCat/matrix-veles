import React, {useState} from "react";
import {useNavigate, useOutlet} from "react-router-dom";

import styles from "./Lists.module.scss";
import {Trans, useTranslation} from "react-i18next";
import ListsTable from "./ListsTable";
import {PreloadedQuery, usePreloadedQuery} from "react-relay/hooks";
import {graphql} from "babel-plugin-relay/macro";
import {ListsQuery} from "./__generated__/ListsQuery.graphql";
import {X} from "lucide-react";
import {LoaderSuspense} from "../../../common/Loader";

type Props = {
    initialQueryRef: PreloadedQuery<ListsQuery>,
}

const Lists = ({initialQueryRef}: Props) => {
    const outlet = useOutlet()
    const navigate = useNavigate()
    const {t} = useTranslation()

    const data = usePreloadedQuery(
            graphql`
                query ListsQuery($first: String, $count: Int) {
                    ...ListsTableFragment
                }
            `,
            initialQueryRef
    )

    const defaultTitle = t("lists.details", "Details")

    const [title, setTitle] = useState(defaultTitle)


    return <div className={styles.roomsContainer}>
        <div className={styles.roomsOverview + (outlet ? " "+styles.leaveSpace : "")}>
            <h1><Trans i18nKey={"lists.title"}>Available Lists</Trans></h1>

            <ListsTable initialQueryRef={data}/>
        </div>

        <div className={styles.slideOver + (outlet ? " "+styles.active : "")}>
            <ListsSlideOverTitleContext.Provider value={setTitle}>
                <div className={styles.slideOverHeader}>
                        <span>{title}</span>
                        <button onClick={() => navigate("/hashing/lists")}><X/></button>
                </div>
                <div className={styles.slideOverContent}>
                    <LoaderSuspense>
                        {outlet}
                    </LoaderSuspense>
                </div>
            </ListsSlideOverTitleContext.Provider>
        </div>
    </div>
}

export const ListsSlideOverTitleContext = React.createContext<React.Dispatch<React.SetStateAction<string>>|null>(null)

export default Lists