import React, { Component } from 'react';
import FileRow from './filerow';

class FileTable extends Component {
    render() {
        const { files, onDelete } = this.props;

        return (
            <div>
                { files.map((file,index) => (
                    <FileRow
                        key={index}
                        id={index}
                        file={file}
                        onDelete={onDelete}
                    />))
                }
            </div>
        );
    }
}

export default FileTable;
