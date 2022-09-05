import React from "react";

import styles from "./ToggleButton.module.scss";

type ToggleButtonProps = {
    name: string
    label: string
    labelSrOnly?: boolean
    onChange?: React.ChangeEventHandler<HTMLInputElement>
    disabled?: boolean
    checked?: boolean
}

const ToggleButton = (props: ToggleButtonProps) => {
    return <label className={styles.toggle} htmlFor={props.name}>
            <input type='checkbox' name={props.name} id={props.name} className={styles.toggleInput} onChange={props.onChange} disabled={props.disabled} checked={props.checked} />
            <span className={styles.toggleDisplay} hidden>
                <span className={styles.spacer}/>I<span className={styles.spacer}/>O<span className={styles.spacer}/>
            </span>
            <span className={props.labelSrOnly ? styles.srOnly : ""}>{props.label}</span>
        </label>;
}

export default ToggleButton