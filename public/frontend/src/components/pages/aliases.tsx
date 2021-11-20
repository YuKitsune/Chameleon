import React, { useState } from 'react';
import useDocTitle from '../hooks/useDocTitle';
import Alias from '../../model/alias';
import AliasList from '../alias/aliasList';
import AddIcon from '../icons/AddIcon';
import Button from '../button';
import Modal, { ModalActionsProps } from '../modal';

const Aliases = () => {

    const [docTitle, setDocTitle] = useDocTitle("Aliases");

		const [isCreateModalOpen, setIsCreateModalOpen] = useState(false);

		let dummyData: Alias[] = [
			{Name:"Facebook", Address: "john.doe@chameleon.io", WhitelistPattern: ".*@facebook\\.com", LastUsed: 5, EncryptionEnabled: true, IsActive: true},
			{Name:"Google", Address: "jupiter_push@chameleon.io", WhitelistPattern: ".*@google\\.com", LastUsed: 3, EncryptionEnabled: true, IsActive: true},
			{Name:"Somewhere Dodgy", Address: "andromeda-audio@chameleon.io", WhitelistPattern: ".*@somewhere-dodgy\\.com", LastUsed: 2, EncryptionEnabled: false, IsActive: false},
		]

    return (
        <div>
						<div className={"flex flex-row justify-between pb-4"}>
            	<h1 className="text-2xl ">Aliases</h1>

							<Button className={"bg-gray-200"} onClick={() => setIsCreateModalOpen(true)}>
								<AddIcon className='h-6 w-6'/>
							</Button>
						</div>
            <AliasList Aliases={dummyData} />
						<Modal
							isOpen={isCreateModalOpen}
							title={"Create Alias"}
							description={"Create an Alias"}
							close={() => setIsCreateModalOpen(false)}
							renderActions={(props: ModalActionsProps) => <Button onClick={props.close}>Close</Button>}>
							Yo
						</Modal>
        </div>
    );
}

export default Aliases;