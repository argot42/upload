import React from 'react';
import styles from './filepicker.module.css';

const FilePicker = (props) => {
    const { onSubmit, onChange } = props;

    return (
        <div className={styles.content}>
            <span className={styles.hiddenInput}>
                <input type="file" onChange={ (e) => {
                    onChange(e.target.files);
                    e.stopPropagation();
                }} multiple />
            </span>
            <br />
            <button className={styles.button} onClick={onSubmit}>upload!</button>
        </div>
    );
}

export default FilePicker;
