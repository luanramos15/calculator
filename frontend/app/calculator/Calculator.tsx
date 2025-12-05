import {Display} from "./components/Display";
import {ButtonPanel, allowedKeys} from "./components/ButtonPanel";
import "./Calculator.css";
import { useState } from "react";
import axios from "axios";

export function Calculator() {
    const [input, setInput] = useState('');

    const handleButtonClick = async(value: string) => {
        if(handleErrors(value)){return;}
        
        if (value === 'Backspace') {
            setInput(input.slice(0, -1));
        } else if (value === 'Escape') {
            setInput('');
        } else if (value === '=' || value === 'Enter') {
            try {
                console.log(input);
                let response = await axios.post('http://localhost:8080/', { expression: input }, {headers: {'Content-Type': 'application/json'}})
                setInput(response.data.result);
                console.log(response);
            } catch (error) {
                setInput('Error');
            }
        } else {
            setInput(input => input + value);
        }
    };

    
    const handleErrors = (value: string) => {
        if (input === 'Error') {setInput(''); return true;}
        if (!allowedKeys.includes(value)){return true;}
        return false;
    }

    const handleKeyPress = (event: object) => {
        const value = event.key;
        handleButtonClick(value);
    }

    return (
        <div
            className="calculator"
            onKeyDown={event => handleKeyPress(event)}
            tabIndex="0"
        >    
            <Display value={input} />
            <ButtonPanel onButtonClick={handleButtonClick} />
        </div>
    );
}