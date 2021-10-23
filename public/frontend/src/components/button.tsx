import React, { PropsWithChildren } from 'react';
import PropsWithClassName from './PropsWithClassName';

type ButtonProps = PropsWithClassName & PropsWithChildren<{}>;

const Button = (props: ButtonProps) => {
    return (
        <div className={`rounded-md px-2 cursor-pointer ${props.className}`}>
            {props.children}
        </div>
    );
};

export default Button;