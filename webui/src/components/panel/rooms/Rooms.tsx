import React from "react";
import {useNavigate, useOutlet} from "react-router-dom";

import styles from "./Rooms.module.scss";
import {Trans} from "react-i18next";
import RoomsTable from "./RoomsTable";
import {PreloadedQuery, usePreloadedQuery, useQueryLoader} from "react-relay/hooks";
import {graphql} from "babel-plugin-relay/macro";
import {RoomsQuery} from "./__generated__/RoomsQuery.graphql";
import {X} from "lucide-react";

type Props = {
    initialQueryRef: PreloadedQuery<RoomsQuery>,
}

const Rooms = ({initialQueryRef}: Props) => {
    const outlet = useOutlet()
    const navigate = useNavigate()

    const data = usePreloadedQuery(
            graphql`
                query RoomsQuery($first: String, $count: Int) {
                    ...RoomsTableFragment
                }
            `,
            initialQueryRef
    )


    return <div className={styles.roomsContainer}>
        <div className={styles.roomsOverview + (outlet ? " "+styles.leaveSpace : "")}>
            <h1><Trans i18nKey={"rooms.title"}>My Rooms</Trans></h1>

            <RoomsTable initialQueryRef={data}/>
        </div>

        <div className={styles.slideOver + (outlet ? " "+styles.active : "")}>
            <div className={styles.slideOverHeader}>
                <span><Trans i18nKey={"rooms.details"}>Details</Trans></span>
                <button onClick={() => navigate("/rooms")}><X/></button>
            </div>
            <div className={styles.slideOverContent}>
                {outlet}
            </div>
        </div>
    </div>
}

export default Rooms