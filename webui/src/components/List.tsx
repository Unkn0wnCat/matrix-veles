import React, {useCallback} from "react";
import styles from "./List.module.scss";
import {Box, Loader} from "lucide-react";
import {Trans, useTranslation} from "react-i18next";
import {Link} from "react-router-dom";
import {LoadMoreFn} from "react-relay/relay-hooks/useLoadMoreFunction";

type Props = {
    children?: React.ReactNodeArray
    isLoadingNext?: boolean
    hasNext?: boolean
    loadNext?: LoadMoreFn<any>|(() => any)
    className?: string
}

const List = (props: Props) => {
    const {t} = useTranslation()

    return <div className={styles.list + (props.className ? " " + props.className :"")}>
        {(!props.children || props.children.length == 0) && <span className={styles.eol}>
                    <Box width={50} height={50} strokeWidth={1} strokeDasharray={"2px 4px"} /><br/>
                    <Trans i18nKey={"list.none"}>No entries</Trans>
                </span>}

        {props.children}

        {props.isLoadingNext && <div className={styles.loader} title={t("list.loading", "Loading more entries...")}><Loader/></div>}
        {!props.isLoadingNext && props.children && props.children.length > 0 && (props.hasNext ? <button className={styles.loadMore} onClick={() => props.loadNext}><Trans i18nKey={"list.more"}>Show more</Trans></button> : <span className={styles.eol}><Trans i18nKey={"list.end"}>No more entries</Trans></span>)}
    </div>
}

export default List