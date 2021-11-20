import React, { PropsWithChildren, useCallback } from 'react';
import PropsWithClassName from './PropsWithClassName';

type ButtonProps = PropsWithClassName & PropsWithChildren<{}> & {
    onClick?: (() => void) | undefined;
    onClickAsync?: (() => Promise<void>) | undefined;
};

const Button = (props: ButtonProps) => {

    const onClick = useCallback(async () => {
        props.onClick && props.onClick();
        props.onClickAsync && await props.onClickAsync();
    }, [props.onClick, props.onClickAsync])

    return (
        <div className={`rounded-md p-1 cursor-pointer ${props.className}`} onClick={onClick}>
            {props.children}
        </div>
    );
};

export default Button;