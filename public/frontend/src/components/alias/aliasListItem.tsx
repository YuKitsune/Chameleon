import React from 'react';
import Alias from '../../model/alias';
import AddIcon from '../icons/AddIcon';
import LockIcon from '../icons/LockIcon';

interface AliasListItemProps {
    Alias?: Alias | undefined;
    IsTemplate?: boolean;
}

const AliasListItem = (props: AliasListItemProps) => {
    const {Alias: alias} = props;

    return (
        <div className='flex justify-between rounded-md bg-gray-100 p-2 flex-row shadow-md cursor-pointer'>

	        <div className='inline-flex grid grid-cols-2 gap-1 flex-grow content-center'>
                {props.IsTemplate && (
                    <>
                        <span className='text-xl row-span-2 text-right'>Add</span>
                        <AddIcon className='row-span-2 h-6 w-6'/>
                    </>
                )}

                {alias && (
                    <>
                        {/* Top row */}
                        <span className='inline-flex gap-2'>
                            <strong>{alias.Name}</strong>
                            {alias.EncryptionEnabled && (
                                <>
                                    {/* Todo: This doesn't quite line up */}
                                    <LockIcon className='stroke-current text-green-500 h-5 w-5' />
                                </>
                            )}
                        </span>
                        <span>Last used: {alias.LastUsed} days ago</span>

                        {/* Bottom row */}
                        <span><pre>{alias.Address}</pre></span>
                    </>
                )}
            </div>
        </div>
    );
}

export default AliasListItem;
