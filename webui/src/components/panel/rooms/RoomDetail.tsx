import React, {useEffect} from "react";
import {PreloadedQuery, usePreloadedQuery} from "react-relay/hooks";
import {graphql} from "babel-plugin-relay/macro";
import {RoomDetailQuery, RoomDetailQuery$variables} from "./__generated__/RoomDetailQuery.graphql";
import {DisposeFn} from "relay-runtime";
import {useParams} from "react-router-dom";
import ToggleButton from "../../form_components/ToggleButton";

import styles from "./RoomDetail.module.scss";
import {useReconfigureRoomMutation} from "../../../mutations/ReconfigureRoomMutation";

type Props = {
    initialQueryRef: PreloadedQuery<RoomDetailQuery> | null | undefined,
    fetch: (variables: RoomDetailQuery$variables) => void
    dispose: DisposeFn
}

type PropsFinal = {
    initialQueryRef: PreloadedQuery<RoomDetailQuery>,
}

const RoomDetailInner = ({initialQueryRef}: PropsFinal) => {
    const [reconfigureRoom, reconfiguringRoom] = useReconfigureRoomMutation();



    const data = usePreloadedQuery(
            graphql`
                query RoomDetailQuery($id: ID) {
                    room(id:$id) {
                        id
                        active
                        deactivated
                        adminPowerLevel
                        debug
                        name
                        roomId
                        hashCheckerConfig {
                            chatNotice
                            hashCheckMode
                            subscribedLists
                        }
                    }
                }
            `,
            initialQueryRef
    )

    return <>
        <span className={styles.title}>
            <h1>{data.room?.name}</h1>
            <ToggleButton name={"activeSwitch"} label={"Activate"} labelSrOnly={true} onChange={(ev) => {
                reconfigureRoom({
                    variables: {
                        reconfigureInput: {
                            id: data.room?.id!,
                            deactivate: !ev.currentTarget.checked
                        }
                    }
                })
            }} disabled={reconfiguringRoom || ((data.room || false) && !data.room.active && !data.room.deactivated)} checked={data.room?.active}/>
        </span>


        <pre>{JSON.stringify(data, null, 2)}</pre>
    </>
}

const RoomDetail = ({initialQueryRef, fetch, dispose}: Props) => {
    const {id} = useParams()

    useEffect(() => {
        fetch({id})

        return () => {
            dispose();
        }
    }, [id, dispose, fetch])

    return initialQueryRef ? <RoomDetailInner initialQueryRef={initialQueryRef} /> : <>loading...</>
}

export default RoomDetail;