import React, { Component } from 'react';
import FileRow from './filerow';
import styles from './filetable.module.css';

class FileTable extends Component {
    render() {
        const { files, onDelete } = this.props;

        return (
            <div className={styles.content}>
                <ul>
                    { files.map((file,index) => (
                        <li key={index}>
                            <FileRow
                              id={index}
                              file={file}
                              onDelete={onDelete}
                            />
                        </li>))
                    }
                </ul>
            </div>
        );
    }
}

export default FileTable;
