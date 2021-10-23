import React from 'react';
import {QuarantinedMail} from '../../model/quarantinedMail';

type QuarantinedMailListItemProps = {
    quarantinedMail: QuarantinedMail;
}

const QuarantinedMailListItem = (props: QuarantinedMailListItemProps) => {
    const {quarantinedMail} = props;

    return (
        <div className='flex justify-between rounded-md bg-gray-100 p-2 flex-row shadow-md cursor-pointer'>

            <div className='inline-flex grid grid-cols-2 gap-1 flex-grow content-center'>

                {/* Top row */}
                <span className='inline-flex gap-1'>From: <strong><pre>{quarantinedMail.Sender}</pre></strong></span>
                <span className='text-right'>{quarantinedMail.DateReceived} days ago</span>

                {/* Middle row */}
                <span className='inline-flex gap-1'>
                    To: <strong className='inline-flex'>{quarantinedMail.Recipient.Name} (<pre>{quarantinedMail.Recipient.Address}</pre>)</strong>
                </span>

                {/* Bottom row */}
                <span className='col-span-2 truncate'>{quarantinedMail.Subject}</span>
            </div>
        </div>
    );
};

export default QuarantinedMailListItem;