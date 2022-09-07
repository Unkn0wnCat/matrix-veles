import React, {useContext, useEffect, useState} from "react";
import {PreloadedQuery, usePreloadedQuery} from "react-relay/hooks";
import {graphql} from "babel-plugin-relay/macro";
import {RoomDetailQuery, RoomDetailQuery$variables} from "./__generated__/RoomDetailQuery.graphql";
import {DisposeFn} from "relay-runtime";
import {useParams} from "react-router-dom";
import {RoomsSlideOverTitleContext} from "./Rooms";
import ToggleButton from "../../form_components/ToggleButton";

import styles from "./RoomDetail.module.scss";
import {useReconfigureRoomMutation} from "../../../mutations/ReconfigureRoomMutation";
import {HashCheckerMode} from "../../../mutations/__generated__/ReconfigureRoomMutation.graphql";
import {Trans, useTranslation} from "react-i18next";

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
    const [newAdminPowerLevel, setNewAdminPowerLevel] = useState<number|null>(null)
    const {t} = useTranslation()


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

    const titleSetContext = useContext(RoomsSlideOverTitleContext)

    titleSetContext && data.room?.name && titleSetContext(data.room.name);

    return <>
        <ToggleButton name={"activeSwitch"} label={t("panel:rooms.detail.activate.label", {defaultValue: "Activate Room"})} labelSrOnly={false} onChange={(ev) => {
            reconfigureRoom({
                variables: {
                    reconfigureInput: {
                        id: data.room?.id!,
                        deactivate: !ev.currentTarget.checked
                    }
                }
            })
        }} disabled={reconfiguringRoom || ((data.room || false) && !data.room.active && !data.room.deactivated)} checked={data.room?.active}/>

        <ToggleButton name={"debugSwitch"} label={t("panel:rooms.detail.debug.label", {defaultValue: "Debug-Mode"})} labelSrOnly={false} onChange={(ev) => {
            reconfigureRoom({
                variables: {
                    reconfigureInput: {
                        id: data.room?.id!,
                        debug: ev.currentTarget.checked
                    }
                }
            })
        }} disabled={reconfiguringRoom || (!data.room)} checked={data.room?.debug}/>

        <label htmlFor={"adminPowerLevelInput"} className={styles.label}><Trans i18nKey={"panel:rooms.detail.adminPowerLevel.label"}>Admin-Power Level</Trans></label>
        <div className={styles.powerLevelInput}>
            <input type={"number"} min={"50"} max={"100"} step={"1"} value={newAdminPowerLevel || data.room?.adminPowerLevel} onChange={(ev) => {
                if(Number.parseInt(ev.currentTarget.value) === data.room?.adminPowerLevel) {
                    setNewAdminPowerLevel(null)
                    return
                }

                setNewAdminPowerLevel(Number.parseInt(ev.currentTarget.value))
            }} id={"adminPowerLevelInput"} />
            <button disabled={reconfiguringRoom || (!newAdminPowerLevel)} onClick={() => {
                if(!newAdminPowerLevel) {
                    return
                }

                reconfigureRoom({
                    variables: {
                        reconfigureInput: {
                            id: data.room?.id!,
                            adminPowerLevel: newAdminPowerLevel
                        }
                    }
                })

                setNewAdminPowerLevel(null)
            }}><Trans i18nKey={"panel:rooms.detail.set.label"}>Set</Trans></button>
        </div>

        <div className={styles.settingsGroup}>
            <span className={styles.settingsGroupTitle}><Trans i18nKey={"panel:rooms.detail.hashChecker.label"}>Hash-Checker</Trans></span>
            <div className={styles.settingsGroupContent}>
                <ToggleButton name={"chatNoticeSwitch"} label={t("panel:rooms.detail.hashChecker.notice.label", {defaultValue: "Send Notice to Chat"})} labelSrOnly={false} onChange={(ev) => {
                    reconfigureRoom({
                        variables: {
                            reconfigureInput: {
                                id: data.room?.id!,
                                hashChecker: {
                                    chatNotice: ev.currentTarget.checked
                                }
                            }
                        }
                    })
                }} disabled={reconfiguringRoom || (!data.room)} checked={data.room?.hashCheckerConfig.chatNotice}/>

                <label htmlFor={"hashCheckModeChooser"} className={styles.label}><Trans i18nKey={"panel:rooms.detail.hashChecker.hashCheckMode.label"}>Behaviour on Match</Trans></label>
                <select value={data.room?.hashCheckerConfig.hashCheckMode} disabled={reconfiguringRoom || (!data.room)} onChange={(ev) => {
                    reconfigureRoom({
                        variables: {
                            reconfigureInput: {
                                id: data.room?.id!,
                                hashChecker: {
                                    hashCheckMode: ev.currentTarget.value as HashCheckerMode
                                }
                            }
                        }
                    })
                }} className={styles.hashCheckModeChooser} id={"hashCheckModeChooser"}>
                    <option value="NOTICE"><Trans i18nKey={"panel:rooms.detail.hashChecker.hashCheckMode.NOTICE.label"}>Only send a Notice</Trans></option>
                    <option value="DELETE"><Trans i18nKey={"panel:rooms.detail.hashChecker.hashCheckMode.DELETE.label"}>Delete the Message</Trans></option>
                    <option value="MUTE"><Trans i18nKey={"panel:rooms.detail.hashChecker.hashCheckMode.MUTE.label"}>Mute User & Delete</Trans></option>
                    <option value="BAN"><Trans i18nKey={"panel:rooms.detail.hashChecker.hashCheckMode.BAN.label"}>Ban User & Delete</Trans></option>
                </select>
            </div>

            <div className={styles.listTable}>
                <table>
                    <thead>
                    <tr>
                        <th><Trans i18nKey={"lists.name"}>Name</Trans></th>
                        <th><Trans i18nKey={"lists.id"}>List ID</Trans></th>
                    </tr>
                    </thead>
                    <tbody>
                    {/*
                        data.rooms?.edges.map((edge) => {
                            return <tr onClick={() => {navigate("/rooms/"+edge.node.id)}} key={edge.node.id}>
                                <td>
                                    {edge.node.debug && <span className={styles.badge + " " + styles.blue}><Trans i18nKey={"rooms.debug"}>Debug</Trans></span>}
                                    {!edge.node.active && <span className={styles.badge + " " + styles.red}><Trans i18nKey={"rooms.inactive"}>Inactive</Trans></span>}
                                    {edge.node.active && <span className={styles.badge + " " + styles.green}><Trans i18nKey={"rooms.active"}>Active</Trans></span>}
                                </td>
                                <td>{edge.node.name}</td>
                                <td>{edge.node.roomId}</td>
                            </tr>;
                        })
                    */}
                    </tbody>
                </table>
            </div>
        </div>


        <pre style={{opacity: .25}}>{JSON.stringify(data, null, 2)}</pre>
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

    return initialQueryRef ? <RoomDetailInner initialQueryRef={initialQueryRef} /> : null
}

export default RoomDetail;