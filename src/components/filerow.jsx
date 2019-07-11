import React, { Component } from 'react';

class FileRow extends Component {
    render() {
        const { file, id, onDelete } = this.props;
        return (
            <div>
                <span>{file.name}</span>
                <button onClick={ () => onDelete(id) }>x</button>
            </div>
        );

    }
}

export default FileRow;
