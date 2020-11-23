import React from 'react';

const userInput = (props) => {
    const inputStyle = {
        border: '1px solid blue'
    };

    return <input 
        type="text" 
        style={inputStyle}
        onChange={props.changed} 
        value={props.currentName} />;
};

export default userInput;