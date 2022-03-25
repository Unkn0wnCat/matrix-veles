import React from "react";
import {graphql} from "babel-plugin-relay/macro";


import {PreloadedQuery, usePreloadedQuery} from "react-relay/hooks";
import {DashboardQuery} from "./__generated__/DashboardQuery.graphql";
import {Trans} from "react-i18next";

type Props = {
    initialQueryRef: PreloadedQuery<DashboardQuery>,
}

const Dashboard = (props: Props) => {
    const data = usePreloadedQuery(
            graphql`
                query DashboardQuery {
                    self {
                        username
                        id
                        admin
                    }
                }
            `,
            props.initialQueryRef
    )

    const name = data.self?.username

    return <>
        <h1><Trans i18nKey={"dashboard.helloText"}>Ayo {{name}}!</Trans></h1>

        {/*<button onClick={refresh}>Refresh</button>*/}

        {/*hasNext && <button
                onClick={() => {
                    loadNext(2)
                }}>
            Load more Entries
        </button>*/}
    </>
}

export default Dashboard