import React from 'react';

const FilePicker = (props) => {
    const { onSubmit, onChange } = props;

    return (
        <div>
            <h1>upload!</h1>
            <form onSubmit={ (e) => { 
                onSubmit();
                e.preventDefault();
            }}>
                <input type="file" onChange={ (e) => {
                    onChange(e.target.files);
                    e.stopPropagation();
                }} multiple />
                <br />
                <button type="submit">Submit</button>
            </form>
        </div>
    );
}

export default FilePicker;
