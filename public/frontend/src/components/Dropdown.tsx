import React, { useState, useRef, useCallback, PropsWithChildren } from 'react';
import useOnClickOutside from './hooks/useClickOutside';

type DropdownProps = PropsWithChildren<{
    button: React.ReactNode;
}>;

const Dropdown = (props: DropdownProps) => {

    const [isOpen, setIsOpen] = useState(false);

    const dropdownRef = useRef<HTMLDivElement>(null);

    const hideDropdown = useCallback(() => setIsOpen(false), []);
    const toggleDropdown = useCallback(() => setIsOpen(!isOpen), [isOpen]);

    useOnClickOutside(dropdownRef, (_: any) => hideDropdown());

    return (
        <div className="relative" ref={dropdownRef}>
            <div onClick={toggleDropdown}>
                {props.button}
            </div>
            {isOpen && (
                <div className='absolute right-0 mt-2 py-2 w-40 bg-white rounded-md shadow-md' onClick={hideDropdown}>
                    {props.children}
                </div>
            )}
        </div>
    );
}

export default Dropdown;
