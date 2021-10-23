import React from 'react';
import Alias from '../../model/alias';
import AliasListItem from './aliasListItem';

interface AliasListProps {
    Aliases: Alias[];
    AllowAdd: boolean;
}

const AliasList = (props: AliasListProps) => {
    return (
        <div className='flex flex-row grid grid-cols-2 gap-4'>
            {props.Aliases.map(alias => {
                return <AliasListItem key={alias.Address} Alias={alias} />
            })}

            {props.AllowAdd && <AliasListItem IsTemplate={true} />}
        </div>
    );
}

export default AliasList;