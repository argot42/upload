import React, { Component } from 'react';
import FilePicker from './components/filepicker';
import FileTable from './components/filetable';
import "./App.css";

class App extends Component {
    state = {
        files: []
    }

    render() {
        return (
            <React.Fragment>
                <FilePicker 
                    onSubmit={this.handleSubmit}
                    onChange={this.handleChange} />
                <FileTable
                    files={this.state.files} 
                    onDelete={this.handleDelete} />
            </React.Fragment>
        );
    }

    handleChange = (selectedFiles) => {
        var newFiles = [];
        for (var i = 0; i < selectedFiles.length; i++) {
            newFiles.push(selectedFiles[i]);
        }
        this.setState({ files: newFiles });
    }

    handleSubmit = () => {
        const data = new FormData();

        for (var file of this.state.files) {
            data.append(file.name, file);
        }

        fetch('/upload', {
            method: 'POST',
            mode: 'same-origin',
            headers: {
                'Content-Type': 'multipart/form-data'
            },
            body: data
        })
        .then(response => {
            alert(response);
        });
    }

    handleDelete = (fileId) => {
        var newFiles = [];
        for (var i = 0; i < this.state.files.length; i++) {
            if (i === fileId) continue;
            newFiles.push(this.state.files[i]);
        }
        this.setState({ files: newFiles });
    }
}

export default App;
