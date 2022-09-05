import React from "react";
import {usePaginationFragment} from "react-relay/hooks";
import {graphql} from "babel-plugin-relay/macro";
import {DashMyRoomsFragment$key} from "./__generated__/DashMyRoomsFragment.graphql";
import {ComponentDashMyRooms} from "./__generated__/ComponentDashMyRooms.graphql";
import {Trans} from "react-i18next";
import styles from "./DashMyRooms.module.scss";
import {Link} from "react-router-dom";
import List from "../../List";

type Props = {
    initialQueryRef: DashMyRoomsFragment$key,
    className?: string,
}

const DashMyRooms = (props: Props) => {
    //const {t} = useTranslation()

    const {data, loadNext, hasNext, isLoadingNext} = usePaginationFragment<ComponentDashMyRooms, DashMyRoomsFragment$key>(
            graphql`
                fragment DashMyRoomsFragment on Query @refetchable(queryName: "ComponentDashMyRooms") {
                    rooms(after: $first, first: $count, filter: {canEdit: true}) @connection(key: "ComponentDashMyRooms_rooms") {
                        edges {
                            node {
                                id
                                name
                                active
                                debug
                                roomId
                            }
                        }
                    }
                }
            `,
            props.initialQueryRef
    )

	return (
		<div className={styles.dashMyRooms + " " + (props.className || "")}>
            <h2><Trans i18nKey={"dashboard.my_rooms.title"}>My Rooms</Trans></h2>

            <List className={styles.list} hasNext={hasNext} isLoadingNext={isLoadingNext} loadNext={loadNext}>
                {
                    data.rooms?.edges.map((edge) => {
                        return <Link className={styles.room} key={edge.node.id} to={"/rooms/"+edge.node.id}>
                            <div className={styles.nameRow}>
                                <span className={styles.name}>{edge.node.name}</span>
                                {edge.node.debug && <span className={styles.badge + " " + styles.blue}><Trans i18nKey={"rooms.debug"}>Debug</Trans></span>}
                                {!edge.node.active && <span className={styles.badge + " " + styles.red}><Trans i18nKey={"rooms.inactive"}>Inactive</Trans></span>}
                                {edge.node.active && <span className={styles.badge + " " + styles.green}><Trans i18nKey={"rooms.active"}>Active</Trans></span>}
                            </div>
                            <span className={styles.id}>{edge.node.roomId}</span>
                        </Link>
                    })
                }
            </List>
		</div>
	)
}

export default DashMyRooms
