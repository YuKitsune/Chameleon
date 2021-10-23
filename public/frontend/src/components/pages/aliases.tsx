import React from 'react';
import useDocTitle from '../hooks/useDocTitle';
import Alias from '../../model/alias';
import AliasList from '../alias/aliasList';

const Aliases = () => {

    const [docTitle, setDocTitle] = useDocTitle("Aliases");

	let dummyData: Alias[] = [
		{Name:"Facebook", Address: "john.doe@chameleon.io", WhitelistPattern: ".*@facebook\\.com", LastUsed: 5, EncryptionEnabled: true, IsActive: true},
		{Name:"Google", Address: "jupiter_push@chameleon.io", WhitelistPattern: ".*@google\\.com", LastUsed: 3, EncryptionEnabled: true, IsActive: true},
		{Name:"Somewhere Dodgy", Address: "andromeda-audio@chameleon.io", WhitelistPattern: ".*@somewhere-dodgy\\.com", LastUsed: 2, EncryptionEnabled: false, IsActive: false}
	]

    return (
        <div>
            <h1 className="text-2xl">Aliases</h1>
            <AliasList Aliases={dummyData} AllowAdd={true}/>
        </div>
    );
}

export default Aliases;