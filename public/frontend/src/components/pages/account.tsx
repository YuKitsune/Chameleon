import React from 'react';
import { TextSize } from '../../model/TextSize';
import Section from '../section';
import TextInput from '../textInput';
import Button from '../button';
import PayPalIcon from '../icons/PayPalIcon';
import ProgressBar from '../progressBar';

const Account = () => {

    // Todo: This needs to be split into a few different components/pages
    // Need to find a better way of organising these since the scree is quite cluttered
    // Using separate pages or tabs might be better here

    return (
        <div className='two-column-grid gap-4'>

            {/* Details */}
            <Section header='My Details'>
                <TextInput label='Primary Email'/>
                <TextInput label='Secondary Email'/>
            </Section>

            {/* Billing */}
            <Section header='Billing'>
                <div className='two-column-grid'>

                    <Section header='Plan' headerSize={TextSize.Large}>
                        <div className='two-column-grid'>

                            {/* Plan title */}
                            <span><strong>Premium +</strong></span>

                            {/* Change plan button */}
                            <div className='inline-flex justify-end'>
                                <Button className='bg-gray-200 hover:bg-gray-300'>Change plan</Button>
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
                                <Button className='bg-gray-200 hover:bg-gray-300'>Change payment method</Button>
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

            {/* Security */}
            <Section header='Security' className='col-span-2'>
                <div className='two-column-grid'>

                    <Section header='Password' headerSize={TextSize.Large}>
                        <TextInput label='Password' isSensitive={true}/>
                        <TextInput label='Confirm Password' isSensitive={true}/>
                    </Section>

                    <Section header='Two Factor Authentication' headerSize={TextSize.Large}>
                        🚧 Todo 🚧
                    </Section>

                    <Section header='Sessions' headerSize={TextSize.Large}>
                        🚧 Todo 🚧
                    </Section>

                    <Section header='Events' headerSize={TextSize.Large}>
                        🚧 Todo 🚧
                    </Section>

                </div>
            </Section>

            {/* Danger Zone */}
            <Section header='Danger Zone' className='col-span-2'>

                <div className='inline-flex'>
                    {/* Delete account */}
                    <Button className='bg-red-500 hover:bg-red-600'>Delete account</Button>
                </div>

            </Section>

        </div>
    );
};

export default Account;