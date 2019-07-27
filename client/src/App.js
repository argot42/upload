import React, { Component } from 'react';
import FilePicker from './components/filepicker';
import FileTable from './components/filetable';
import ProgressBar from './components/progressbar';
import "./App.css";

class App extends Component {
    state = {
        files: [],
        percentComplete: 0,
    }

    render() {
        return (
            <React.Fragment>
                <ProgressBar 
                    percentage={this.state.percentComplete} />
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

        var XHR = new XMLHttpRequest();
        XHR.open('POST', '/upload', true); // async = true
        XHR.onreadystatechange = () => {
            if (XHR.readyState === 4 && XHR.status === 200) {
                this.setState({ files: [] });
            }
        }
        XHR.upload.addEventListener('progress', (e) => {
            if (e.lengthComputable) {
                var percentComplete = Math.round(e.loaded * 100 / e.total);
                this.setState({ percentComplete: percentComplete });
            }
        }, false);
        XHR.send(data);
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
