import React from 'react';
import styles from './progressbar.module.css';

const ProgressBar = (props) => {
    return (
        <div className={styles.container}>
            <div style={{ width: props.percentage.toString() + '%' }} className={styles.bar} />
        </div>
    );
}

export default ProgressBar;
