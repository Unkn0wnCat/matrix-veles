import React, {useCallback, useEffect, useRef, useState} from "react";
import {useNavigate, useOutlet} from "react-router-dom";

import styles from "./Entries.module.scss";
import {Trans, useTranslation} from "react-i18next";
import EntriesTable from "./EntriesTable";
import {PreloadedQuery, usePreloadedQuery} from "react-relay/hooks";
import {graphql} from "babel-plugin-relay/macro";
import {EntriesQuery} from "./__generated__/EntriesQuery.graphql";
import {X} from "lucide-react";
import {LoaderSuspense} from "../../../common/Loader";
import {sha512} from "js-sha512";
import {ReactComponent as Logo} from "../../../../logo.svg";
import anime, {AnimeInstance} from "animejs";

type Props = {
    initialQueryRef: PreloadedQuery<EntriesQuery>,
}

type UploadQueueEntry = {
    fileName: string
    fileHash: string
}

const Entries = ({initialQueryRef}: Props) => {
    const outlet = useOutlet()
    const navigate = useNavigate()
    const {t} = useTranslation()

    const data = usePreloadedQuery(
            graphql`
                query EntriesQuery($first: String, $count: Int) {
                    ...EntriesTableFragment
                }
            `,
            initialQueryRef
    )

    const defaultTitle = t("entries.details", "Details")

    const [title, setTitle] = useState(defaultTitle)

    const [showUploader, setShowUploader] = useState(false)
    const [uploadCounter, setUploadCounter] = useState(0)

    const uploadQueue = useRef<UploadQueueEntry[]>([])
    const incompleteCounter = useRef(0)
    const loadingAnimation = useRef<AnimeInstance>()
    const loadingLogo = useRef<SVGSVGElement>(null)

    const setupAnim = useCallback(() => {
            if(!loadingLogo.current) return

            if(incompleteCounter.current <= 0) {
                anime({
                    targets: "#"+loadingLogo.current.id+" .number",
                    duration: 500,
                    easing: "easeInOutCubic",
                    opacity: 1,
                })
                loadingAnimation.current = undefined
                return
            }

            loadingAnimation.current = anime({
                targets: "#"+loadingLogo.current.id+" .number",
                duration: 500,
                easing: "easeInOutCubic",
                opacity: () => anime.random(.25, 1),
                complete: () => setupAnim()
            })
        }, [])

    useEffect(() => {
        if(!loadingLogo.current) return

        if(!loadingAnimation.current && incompleteCounter.current > 0) setupAnim()
    }, [uploadCounter, setupAnim])

    const processFile = (file: File) => {

        const reader = new FileReader()

        reader.addEventListener('load', (ev) => {
            const res = ev.target!.result as ArrayBuffer;

            const hash = sha512(res)

            uploadQueue.current.push({
                fileName: file.name,
                fileHash: hash,
            })

            incompleteCounter.current--

            setUploadCounter(prev => prev + 1)
        })

        reader.addEventListener('error', (err) => {
            console.error(err)
            incompleteCounter.current--
        })

        reader.readAsArrayBuffer(file)
    }

    const processDirectory = (directory: FileSystemDirectoryEntry) => {
        const reader = directory.createReader()
        reader.readEntries((entries) => {
            entries.forEach((entry) => {
                if(entry.isFile) {
                    (entry as FileSystemFileEntry).file((file) => {
                        processFile(file)
                        incompleteCounter.current++
                    })
                    return
                }

                if(entry.isDirectory) {
                    processDirectory(entry as FileSystemDirectoryEntry)
                    incompleteCounter.current++
                    return
                }
            })
            incompleteCounter.current--
        }, (err) => {
            console.error("Could not read directory", directory.fullPath, err)
            incompleteCounter.current--
        })
    }

    const processDataTransfer = (items: DataTransferItemList) => {

        for(let i = 0; i < items.length; i++) {
            const item = items[i];

            const entry = item.webkitGetAsEntry()

            if(!entry) {
                continue
            }

            if(entry.isDirectory) {
                const dir = entry as FileSystemDirectoryEntry;

                processDirectory(dir)
                continue
            }
            if(entry.isFile) {
                const file = entry as FileSystemFileEntry;

                file.file((file) => {
                    processFile(file)
                })
                continue
            }
        }
    }

    return <div className={styles.roomsContainer}>
        <div className={styles.roomsOverview + (outlet ? " "+styles.leaveSpace : "")}>
            <h1><Trans i18nKey={"entries.title"}>Entry Management</Trans></h1>
            <button onClick={() => setShowUploader(true)}>Upload New</button>

            <EntriesTable initialQueryRef={data}/>
        </div>

        <div className={styles.modalOuter + (showUploader ? " "+styles.active :"")}>
            <div className={styles.uploader} role={"dialog"} aria-modal={true}>
                <span className={styles.modalTitle}>Upload entries</span>

                <p>In this form you'll be able to upload new entries. Only the hashes of the files will be sent to the server. The actual files never leave your device.</p>

                {/*<div>
                    <input type={"text"} placeholder={"Lists to upload to"} />
                </div>*/}

                <div className={styles.dropArea} onDragOver={(ev) => {
                    ev.stopPropagation();
                    ev.preventDefault();
                    ev.dataTransfer.dropEffect = 'copy';
                }} onDrop={(ev) => {
                    ev.stopPropagation();
                    ev.preventDefault();
                    const fileList = ev.dataTransfer.items;
                    processDataTransfer(fileList)
                }}>
                    <div className={styles.dropAreaContent}>
                        <Logo ref={loadingLogo} id={"loadingLogo"}/>
                        <span>{incompleteCounter.current > 0 ? "Hashing..." : "Drop files to hash"}</span>
                    </div>
                </div>

                <div className={styles.fileList}>
                    {
                        uploadQueue.current.reverse().slice(0, 5).map((entry) => {
                            return <div className={styles.file}>
                                <span className={styles.fileName}>{entry.fileName}</span>
                                <span className={styles.fileHash}>{entry.fileHash}</span>
                            </div>
                        })
                    }
                    {uploadQueue.current.length > 5 && <span>And {uploadQueue.current.length - 5} more...</span>}
                </div>

                <button onClick={() => setShowUploader(false)}>Cancel</button>
            </div>
        </div>

        <div className={styles.slideOver + (outlet ? " "+styles.active : "")}>
            <EntriesSlideOverTitleContext.Provider value={setTitle}>
                <div className={styles.slideOverHeader}>
                        <span>{title}</span>
                        <button onClick={() => navigate("/hashing/entries")}><X/></button>
                </div>
                <div className={styles.slideOverContent}>
                    <LoaderSuspense>
                        {outlet}
                    </LoaderSuspense>
                </div>
            </EntriesSlideOverTitleContext.Provider>
        </div>
    </div>
}

export const EntriesSlideOverTitleContext = React.createContext<React.Dispatch<React.SetStateAction<string>>|null>(null)

export default Entries