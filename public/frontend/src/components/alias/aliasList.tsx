import React from 'react';
import Alias from '../../model/alias';
import AliasListItem from './aliasListItem';

interface AliasListProps {
    Aliases: Alias[];
}

const AliasList = (props: AliasListProps) => {
    return (
        <div className='flex flex-row grid grid-cols-1 gap-4'>
            {props.Aliases.map(alias => {
                return <AliasListItem key={alias.Address} Alias={alias} />
            })}
        </div>
    );
}

export default AliasList;