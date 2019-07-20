import React, { Component } from 'react';
import styles from './filerow.module.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faWindowClose } from '@fortawesome/free-solid-svg-icons';

class FileRow extends Component {
    render() {
        const { file, id, onDelete } = this.props;
        return (
            <div>
                <FontAwesomeIcon className={styles.icon} size="sm" onClick={ () => onDelete(id) } icon={faWindowClose} />
                <span className={styles.container}>
                    {file.name}
                </span>
            </div>
        );

    }
}

export default FileRow;
