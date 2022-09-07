import React from "react";
import {usePaginationFragment} from "react-relay/hooks";
import {graphql} from "babel-plugin-relay/macro";
import {ListsTableFragment$key} from "./__generated__/ListsTableFragment.graphql";
import styles from "./ListsTable.module.scss";
import {useNavigate} from "react-router-dom";
import {Trans} from "react-i18next";

type Props = {
    initialQueryRef: ListsTableFragment$key,
}

const ListsTable = ({initialQueryRef}: Props) => {
    const {data} = usePaginationFragment(graphql`
        fragment ListsTableFragment on Query @refetchable(queryName: "ListsTableFragment") {
            lists(after: $first, first: $count) @connection(key: "ListsTableFragment_lists") {
                edges {
                    node {
                        id
                        name
                        tags
                        creator {
                            id
                            username
                            matrixLinks
                        }
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
                    <th><Trans i18nKey={"lists.name"}>Name</Trans></th>
                    <th><Trans i18nKey={"lists.id"}>List ID</Trans></th>
                    <th><Trans i18nKey={"lists.creator"}>List Creator</Trans></th>
                </tr>
            </thead>
            <tbody>
            {
                data.lists?.edges.map((edge) => {
                    return <tr onClick={() => {navigate("/hashing/lists/"+edge.node.id)}} key={edge.node.id}>
                        <td>{edge.node.name}</td>
                        <td>{edge.node.id}</td>
                        <td>{edge.node.creator.username}</td>
                    </tr>;
                })
            }
            </tbody>

        </table>
    </div>
}

export default ListsTable