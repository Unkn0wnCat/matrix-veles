import React, {useEffect} from "react";
import {PreloadedQuery, usePreloadedQuery} from "react-relay/hooks";
import {graphql} from "babel-plugin-relay/macro";
import {RoomDetailQuery, RoomDetailQuery$variables} from "./__generated__/RoomDetailQuery.graphql";
import {DisposeFn} from "relay-runtime";
import {useParams} from "react-router-dom";

type Props = {
    initialQueryRef: PreloadedQuery<RoomDetailQuery> | null | undefined,
    fetch: (variables: RoomDetailQuery$variables) => void
    dispose: DisposeFn
}

type PropsFinal = {
    initialQueryRef: PreloadedQuery<RoomDetailQuery>,
}

const RoomDetailInner = ({initialQueryRef}: PropsFinal) => {
    const data = usePreloadedQuery(
            graphql`
                query RoomDetailQuery($id: ID) {
                    room(id:$id) {
                        id
                        active
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
        <h1>Room Detail: {data.room?.name}</h1>

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