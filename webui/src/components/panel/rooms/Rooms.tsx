import React, {useState} from "react";
import {useNavigate, useOutlet} from "react-router-dom";

import styles from "./Rooms.module.scss";
import {Trans, useTranslation} from "react-i18next";
import RoomsTable from "./RoomsTable";
import {PreloadedQuery, usePreloadedQuery} from "react-relay/hooks";
import {graphql} from "babel-plugin-relay/macro";
import {RoomsQuery} from "./__generated__/RoomsQuery.graphql";
import {X} from "lucide-react";
import {LoaderSuspense} from "../../common/Loader";

type Props = {
    initialQueryRef: PreloadedQuery<RoomsQuery>,
}

const Rooms = ({initialQueryRef}: Props) => {
    const outlet = useOutlet()
    const navigate = useNavigate()
    const {t} = useTranslation()

    const data = usePreloadedQuery(
            graphql`
                query RoomsQuery($first: String, $count: Int) {
                    ...RoomsTableFragment
                }
            `,
            initialQueryRef
    )

    const defaultTitle = t("rooms.details", "Details")

    const [title, setTitle] = useState(defaultTitle)


    return <div className={styles.roomsContainer}>
        <div className={styles.roomsOverview + (outlet ? " "+styles.leaveSpace : "")}>
            <h1><Trans i18nKey={"rooms.title"}>My Rooms</Trans></h1>

            <RoomsTable initialQueryRef={data}/>
        </div>

        <div className={styles.slideOver + (outlet ? " "+styles.active : "")}>
            <RoomsSlideOverTitleContext.Provider value={setTitle}>
                <div className={styles.slideOverHeader}>
                        <span>{title}</span>
                        <button onClick={() => navigate("/rooms")}><X/></button>
                </div>
                <div className={styles.slideOverContent}>
                    <LoaderSuspense>
                        {outlet}
                    </LoaderSuspense>
                </div>
            </RoomsSlideOverTitleContext.Provider>
        </div>
    </div>
}

export const RoomsSlideOverTitleContext = React.createContext<React.Dispatch<React.SetStateAction<string>>|null>(null)

export default Rooms