import React from "react";
import {usePaginationFragment} from "react-relay/hooks";
import {graphql} from "babel-plugin-relay/macro";
import {Trans} from "react-i18next";
import styles from "./DashMyLists.module.scss";
import {Link} from "react-router-dom";
import List from "../../List";
import {DashMyListsFragment$key} from "./__generated__/DashMyListsFragment.graphql";
import {ComponentDashMyLists} from "./__generated__/ComponentDashMyLists.graphql";

type Props = {
    initialQueryRef: DashMyListsFragment$key,
    className?: string,
}

const DashMyLists = (props: Props) => {
    //const {t} = useTranslation()

    const {data, loadNext, hasNext, isLoadingNext} = usePaginationFragment<ComponentDashMyLists, DashMyListsFragment$key>(
            graphql`
                fragment DashMyListsFragment on Query @refetchable(queryName: "ComponentDashMyLists") {
                    lists(after: $first, first: $count) @connection(key: "ComponentDashMyLists_lists") {
                        edges {
                            node {
                                id
                                name
                            }
                        }
                    }
                }
            `,
            props.initialQueryRef
    )

	return (
		<div className={styles.dashMyLists + " " + (props.className || "")}>
            <h2><Trans i18nKey={"dashboard.my_lists.title"}>My Lists</Trans></h2>

            <List className={styles.list} hasNext={hasNext} isLoadingNext={isLoadingNext} loadNext={loadNext}>
                {
                    data.lists?.edges.map((edge) => {
                        return <Link className={styles.listEntry} key={edge.node.id} to={"/hashing/lists/"+edge.node.id}>
                            <div className={styles.nameRow}>
                                <span className={styles.name}>{edge.node.name}</span>
                            </div>
                            <span className={styles.id}>{edge.node.id}</span>
                        </Link>
                    })
                }
            </List>
		</div>
	)
}

export default DashMyLists
