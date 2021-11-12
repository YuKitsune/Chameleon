import React, { PropsWithChildren, useCallback, useState } from 'react';
import { Transition } from '@headlessui/react';
import Button from './button';
import ClipboardIcon from './icons/ClipboardIcon';
import ClipboardCheckIcon from './icons/ClipboardCheckIcon';

type CopyToClipboardProps = PropsWithChildren<{}> & {
	text: string;
	activeTimeoutSeconds?: number | undefined;
};

const CopyToClipboardButton = (props: CopyToClipboardProps) => {
	const [isActive, setIsActive] = useState(false);

	let activeTimeoutMs = 3 * 1000;
	if (props.activeTimeoutSeconds)
		activeTimeoutMs = props.activeTimeoutSeconds * 1000;

	const cb = useCallback(
		async () => {
			await navigator.clipboard.writeText(props.text);
			setIsActive(true);

			setTimeout(
				() => setIsActive(false),
				activeTimeoutMs
			);
		},
		[props.text]
	);

	return (
		<Button onClick={cb}>

			{/* Inactive button */}
			{!isActive && <ClipboardIcon className={"h-4, w-4"} />}

			{/* Active button */}
			{isActive && <ClipboardCheckIcon className={"h-4, w-4 text-green-500"} />}

		</Button>
	)
}

export default CopyToClipboardButton;