import React, {useCallback} from "react";
import {graphql} from "babel-plugin-relay/macro";


import {PreloadedQuery, usePaginationFragment, usePreloadedQuery} from "react-relay/hooks";
import {DashboardQuery} from "./__generated__/DashboardQuery.graphql";
import {Trans} from "react-i18next";
import {DashboardListsQuery} from "./__generated__/DashboardListsQuery.graphql";
import {DashboardQueryLists$data, DashboardQueryLists$key} from "./__generated__/DashboardQueryLists.graphql";

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
                    ...DashboardQueryLists
                }
            `,
            props.initialQueryRef
    )

    const {data: d2, hasNext, loadNext, refetch} = usePaginationFragment<DashboardListsQuery, DashboardQueryLists$key>(
            graphql`
                fragment DashboardQueryLists on Query @refetchable(queryName: "DashboardListsQuery") {
                    rooms(after: $first, first: $count, filter: {canEdit: true}) @connection(key: "DashboardQueryLists_rooms") {
                        edges {
                            node {
                                id
                                active
                                debug
                                hashCheckerConfig {
                                    chatNotice
                                    hashCheckMode
                                }
                            }
                        }
                    }
                }
            `,
            // @ts-ignore
            data
    )

    const refresh = useCallback(() => {
        refetch({}, {fetchPolicy: "network-only"})
    }, [])


    const name = data.self?.username

    return <>
        <h1><Trans i18nKey={"dashboard.helloText"}>Ayo {{name}}!</Trans></h1>

        {<button onClick={refresh}>Refresh</button>}

        <pre>{
            JSON.stringify(d2, null, 2)
        }</pre>

        {hasNext && <button
                onClick={() => {
                    loadNext(2)
                }}>
            Load more Entries
        </button>}
    </>
}

export default Dashboard