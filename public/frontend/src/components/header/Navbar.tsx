import React from 'react';
import Link from './Link';
import NotificationIcon from '../icons/NotificationIcon';
import Dropdown from '../Dropdown';
import DropdownItem from '../DropdownItem';
import UserCircle from '../icons/UserCircle';
import '../DropdownItem.css';

const Navbar = () => {
    return (
        <div className="items-center justify-between flex bg-green-500 text-white bg-opacity-90 px-24 py-2 shadow-2xl">

            {/* Left */}
            <div className="inline-flex items-center">

                {/* Todo: Logo */}
                <div className="text-2xl text-white font-semibold inline-flex">
                    <span>Chameleon</span>
                </div>

                {/* Links */}
                <div className='ml-5 flex'>
                    <Link className='mx-2' path='/aliases'>Aliases</Link>
                    <Link className='mx-1' path='/quarantined'>Quarantined</Link>
                </div>
            </div>

            {/* Right */}
            {/* Todo: These should be dropdowns */}
            <div className="inline-flex items-center">
                <Link className='mx-1' path='/notifications'><NotificationIcon HasUnreadNotifications={false} className="h-6 w-6"/></Link>

                {/* Todo: The active state of these links looks weird */}
                <Dropdown button={<Link className='mx-1'><UserCircle className="h-6 w-6"/></Link>}>
                    <DropdownItem>
                        <Link path='/account'>My Account</Link>
                    </DropdownItem>
                    <DropdownItem>
                        <Link path='/settings'>Settings</Link>
                    </DropdownItem>
                    <DropdownItem>
                        <Link path='/logout'>Log Out</Link>
                    </DropdownItem>
                </Dropdown>
            </div>
        </div>
    );
}

export default Navbar;