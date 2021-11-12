import React from 'react';
import Section from '../section';
import { TextSize } from '../../model/TextSize';
import Button from '../button';
import PayPalIcon from '../icons/PayPalIcon';
import ProgressBar from '../progressBar';

const Billing = () => {
	return (
		<Section header='Billing'>
			<div className='two-column-grid'>

				<Section header='Plan' headerSize={TextSize.Large}>
					<div className='two-column-grid'>

						{/* Plan title */}
						<span><strong>Premium +</strong></span>

						{/* Change plan button */}
						<div className='inline-flex justify-end'>
							<Button className='bg-gray-200 hover:bg-gray-300' onClick={async () => {}}>Change plan</Button>
						</div>

						{/* Next billing date */}
						<span className='col-span-2'>Next billing date: 1/12/2022</span>

						<div>
							<label htmlFor='auto-renew-checkbox'>Auto-renew</label>
							<input type='checkbox' id='auto-renew-checkbox'/>
						</div>
					</div>
				</Section>

				<Section header='Payment Method' headerSize={TextSize.Large}>
					<div className='one-column-grid'>

						{/* Selected payment method */}
						<div className='flex inline-flex align-middle'>
							<PayPalIcon className='w-16 mt-1 mr-2' /> <strong>****doe@mymail.com</strong>
						</div>

						{/* Change Payment Method button */}
						{/* Todo: Let's not use a button here... It's kinda wide */}
						<div className='inline-flex justify-end'>
							<Button className='bg-gray-200 hover:bg-gray-300' onClick={async () => {}}>Change payment method</Button>
						</div>
					</div>
				</Section>

				{/* Usage */}
				<Section header='Usage' headerSize={TextSize.Large} className='col-span-2'>
					<div className='two-column-grid'>
						<div className='flex flex-col'>
							<strong>Cached Mail</strong>
							<ProgressBar value={150} max={1024} className='h-3 m-1' />
						</div>
						<div className='flex flex-col'>
							<strong>Quarantined Mail</strong>
							<ProgressBar value={256} max={1024} className='h-3 m-1' />
						</div>
					</div>
				</Section>

			</div>
		</Section>
	);
}

export default Billing;