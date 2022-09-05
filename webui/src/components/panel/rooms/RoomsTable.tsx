import React from "react";
import {usePaginationFragment} from "react-relay/hooks";
import {graphql} from "babel-plugin-relay/macro";
import {RoomsTableFragment$key} from "./__generated__/RoomsTableFragment.graphql";
import styles from "./RoomsTable.module.scss";
import {useNavigate} from "react-router-dom";
import {Trans} from "react-i18next";

type Props = {
    initialQueryRef: RoomsTableFragment$key,
}

const RoomsTable = ({initialQueryRef}: Props) => {
    const {data} = usePaginationFragment(graphql`
        fragment RoomsTableFragment on Query @refetchable(queryName: "RoomsTableFragment") {
            rooms(after: $first, first: $count, filter: {canEdit: true}) @connection(key: "RoomsTableFragment_rooms") {
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
    `, initialQueryRef)

    const navigate = useNavigate()

    return <div className={styles.roomsTableWrapper}>
        <table className={styles.roomsTable}>
            <thead>
                <tr>
                    <th></th>
                    <th><Trans i18nKey={"rooms.name"}>Name</Trans></th>
                    <th><Trans i18nKey={"rooms.id"}>Room ID</Trans></th>
                </tr>
            </thead>
            <tbody>
            {
                data.rooms?.edges.map((edge) => {
                    return <tr onClick={() => {navigate("/rooms/"+edge.node.id)}}>
                        <td>
                            {edge.node.debug && <span className={styles.badge + " " + styles.blue}><Trans i18nKey={"rooms.debug"}>Debug</Trans></span>}
                            {!edge.node.active && <span className={styles.badge + " " + styles.red}><Trans i18nKey={"rooms.inactive"}>Inactive</Trans></span>}
                            {edge.node.active && <span className={styles.badge + " " + styles.green}><Trans i18nKey={"rooms.active"}>Active</Trans></span>}
                        </td>
                        <td>{edge.node.name}</td>
                        <td>{edge.node.roomId}</td>
                    </tr>;
                })
            }
            </tbody>

        </table>
    </div>
}

export default RoomsTable