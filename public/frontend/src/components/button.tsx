import React, { PropsWithChildren } from 'react';
import PropsWithClassName from './PropsWithClassName';

type ButtonProps = PropsWithClassName & PropsWithChildren<{}> & {
    onClick: () => Promise<void>;
};

const Button = (props: ButtonProps) => {
    return (
        <div className={`rounded-md p-1 cursor-pointer ${props.className}`} onClick={props.onClick}>
            {props.children}
        </div>
    );
};

export default Button;