
export const buttons = [
    '7', '8', '9', '/',
    '4', '5', '6', '*',
    '1', '2', '3', '-',
    '0', '.', '=', '+',
    '(', ')', '^', '√'
];

export const allowedKeys = [
    '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
    '/', '*', '-', '+', '.', '(', ')', '^', '√', 'Enter', '=', 'Backspace', 'Escape'
];

export function ButtonPanel({ onButtonClick }) {

    return (
        <div className="buttons">
            <div className="button-ac">
                <button key={'A/C'} onClick={() => onButtonClick('Escape')}>
                        {'A/C'}
                </button>
            </div>
            <div className="button-panel">
                {buttons.map((btn) => (
                    <button key={btn} onClick={() => onButtonClick(btn)}>
                        {btn}
                    </button>
                ))}
            </div>
        </div>
    );
}