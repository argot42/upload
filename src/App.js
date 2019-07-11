import React, { Component } from 'react';
import FilePicker from './components/filepicker';
import FileTable from './components/filetable';

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
        console.log(this.state.files);
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
