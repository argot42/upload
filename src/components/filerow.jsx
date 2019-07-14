import React, { Component } from 'react';
import styles from './filerow.module.css';

class FileRow extends Component {
    render() {
        const { file, id, onDelete } = this.props;
        return (
            <div className={styles.content}>
                <span onClick={ () => onDelete(id) }>{file.name}</span>
            </div>
        );

    }
}

export default FileRow;
