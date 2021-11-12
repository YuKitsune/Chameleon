import React from 'react';
import { TextSize } from '../../model/TextSize';
import Section from '../section';
import TextInput from '../textInput';
import Button from '../button';

const Account = () => {

    // Todo: This needs to be split into a few different components/pages
    // Need to find a better way of organising these since the scree is quite cluttered
    // Using separate pages or tabs might be better here

    return (
      <Section header='My Account'>
        <div className='flex flex-col gap-4'>

            {/* Details */}
            <Section header='My Details' headerSize={TextSize.Large}>
                <TextInput label='Primary Email'/>
                <TextInput label='Secondary Email'/>
            </Section>

            {/* Security */}
            <Section header='Security' className='col-span-2' headerSize={TextSize.Large}>
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
            <Section header='Danger Zone' className='col-span-2' headerSize={TextSize.Large}>

                <div className='inline-flex'>
                    {/* Delete account */}
                    <Button className='bg-red-500 hover:bg-red-600 text-white' onClick={async () => {}}>Delete account</Button>
                </div>

            </Section>

        </div>
      </Section>
    );
};

export default Account;