import React from 'react';
import {QuarantinedMail} from '../../model/quarantinedMail';
import QuarantinedMailListItem from './quarantinedMailListItem';

type QuarantinedMailListProps = {
    quarantinedMail: QuarantinedMail[];
}

const QuarantinedMailList = (props: QuarantinedMailListProps) => {
    const {quarantinedMail} = props;

    // <!-- Todo: This and the ListItem need some re-work -->
    return (
        <div className='flex flex-col grid grid-cols-2 gap-4'>
            {quarantinedMail.map(item =>
                <QuarantinedMailListItem key={item.Sender} quarantinedMail={item} />
                )}
        </div>
    );
};

export default QuarantinedMailList;