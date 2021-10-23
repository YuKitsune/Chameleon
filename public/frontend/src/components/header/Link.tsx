import React, {PropsWithChildren} from 'react';
import { NavLink } from 'react-router-dom';
import PropsWithClassName from '../PropsWithClassName';
import './Link.css';

type LinkProps = PropsWithChildren<{
    path?: string | undefined;
}> & PropsWithClassName;

const Link = (props: LinkProps) => {

    let linkClassName = `link ${props.className}`;

    return (
        <>
            {props.path && (
                <NavLink to={props.path} className={linkClassName} activeClassName={"active"} >
                    {props.children}
                </NavLink>
            )}
            {!props.path && (
                <a className={linkClassName} >
                    {props.children}
                </a>
            )}
        </>
    )
}

export default Link;
