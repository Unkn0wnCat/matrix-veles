import React from "react";
import {graphql} from "babel-plugin-relay/macro";


import {PreloadedQuery, usePreloadedQuery} from "react-relay/hooks";
import {DashboardQuery} from "./__generated__/DashboardQuery.graphql";
import {Trans} from "react-i18next";
import DashMyRooms from "./DashMyRooms";
import styles from "./Dashboard.module.scss";
import DashMyLists from "./DashMyLists";

type Props = {
    initialQueryRef: PreloadedQuery<DashboardQuery>,
}

const Dashboard = (props: Props) => {
    const data = usePreloadedQuery<DashboardQuery>(
            graphql`
                query DashboardQuery($first: String, $count: Int) {
                    self {
                        username
                        id
                        admin
                    }
                    ...DashMyRoomsFragment
                    ...DashMyListsFragment
                }
            `,
            props.initialQueryRef
    )


    const name = data.self?.username

    return <>
        <h1><Trans i18nKey={"dashboard.helloText"}>Ayo {{name}}!</Trans></h1>

        <div className={styles.dashboardGrid}>
            <DashMyRooms initialQueryRef={data}/>
            <DashMyLists initialQueryRef={data}/>
        </div>

    </>
}

export default Dashboard