import React from "react";
import {usePaginationFragment} from "react-relay/hooks";
import {graphql} from "babel-plugin-relay/macro";
import {EntriesTableFragment$key} from "./__generated__/EntriesTableFragment.graphql";
import styles from "./EntriesTable.module.scss";
import {useNavigate} from "react-router-dom";
import {Trans} from "react-i18next";

type Props = {
    initialQueryRef: EntriesTableFragment$key,
}

const EntriesTable = ({initialQueryRef}: Props) => {
    const {data} = usePaginationFragment(graphql`
        fragment EntriesTableFragment on Query @refetchable(queryName: "EntriesTableFragment") {
            entries(after: $first, first: $count) @connection(key: "EntriesTableFragment_entries") {
                edges {
                    node {
                        id
                        tags
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
                    <th><Trans i18nKey={"entries.id"}>Entry ID</Trans></th>
                </tr>
            </thead>
            <tbody>
            {
                data.entries?.edges.map((edge) => {
                    return <tr onClick={() => {navigate("/hashing/entries/"+edge.node.id)}} key={edge.node.id}>
                        <td>{edge.node.id}</td>
                    </tr>;
                })
            }
            </tbody>

        </table>
    </div>
}

export default EntriesTable