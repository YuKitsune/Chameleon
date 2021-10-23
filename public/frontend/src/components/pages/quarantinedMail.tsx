import React from 'react';
import {QuarantinedMail as Model} from '../../model/quarantinedMail';
import useDocTitle from '../hooks/useDocTitle';
import QuarantinedMailList from '../quarantined/quarantinedMailList';

const QuarantinedMail = () => {

    const [docTitle, setDocTitle] = useDocTitle("Quarantined Mail");

	let dummyData: Model[] = [
		{Sender:"danny.devito@macaroni.cheese", Recipient: {Name:"Facebook", Address: "john.doe@chameleon.io"}, Subject: "Smell my nuts", DateReceived: 5, VirusTotalRating: 1, DomainTrustRating: 1},
		{Sender:"sammy.von.soomel@nintento.info", Recipient: {Name:"Facebook", Address: "john.doe@chameleon.io"}, Subject: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", DateReceived: 2, VirusTotalRating: 1, DomainTrustRating: 1},
	]

    return (
        <div>
            <h1 className='text-2xl'>Quarantined</h1>
            <QuarantinedMailList quarantinedMail={dummyData} />
        </div>
    );
};

export default QuarantinedMail;