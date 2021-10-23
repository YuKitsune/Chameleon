import React from 'react';
import PropsWithClassName from './PropsWithClassName';

type TextInputProps = PropsWithClassName & {
    label: string;
    placeholder?: string;
    name?: string;
    isSensitive?: boolean;
};

const TextInput = (props: TextInputProps) => {

    let type = props.isSensitive ? "password" : "text";

    return (
        <div className='py-1 {className}'>
            <label className="block text-gray-700 text-sm font-bold mb-2" htmlFor={props.name}>
                {props.label}
            </label>
            <input className="shadow appearance-none border rounded-sm w-full py-1 px-2 text-gray-700 leading-tight focus:outline-none focus:shadow-outline focus:border-green-500" type={type} placeholder={props.placeholder} id={props.name} />
        </div>
    );
};

export default TextInput;