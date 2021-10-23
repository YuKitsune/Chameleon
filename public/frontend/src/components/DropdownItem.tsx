import React, { PropsWithChildren } from 'react';
import './DropdownItem.css';

type DropdownItemProps = PropsWithChildren<{}>;

const DropdownItem = (props: DropdownItemProps) => {
    return (
        <div className="dropdown-item">
            {props.children}
        </div>
    );
}

export default DropdownItem;