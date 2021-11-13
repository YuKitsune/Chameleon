import React from 'react';
import Section from '../section';
import { TextSize } from '../../model/TextSize';
import Listbox from '../Listbox';

const Preferences = () => {

    return (
        <Section header='Preferences'>
            <div className='flex flex-col gap-4'>

                {/* Forgive me future me, for this scuffness */}

                <div className={"inline flex items-center"}>
                    <span>When a sender doesn’t match the whitelist pattern</span>
                    <Listbox className={"ml-5"} items={[
                      { Label: "Move to quarantine", Value: "quarantine", Disabled: false},
                      { Label: "Delete", Value: "delete", Disabled: false},
                    ]}/>
                </div>

              <div className={"inline flex items-center"}>
                <label htmlFor='notify-when-quarantined-checkbox'>Notify me when a new message has been quarantined</label>
                <input type='checkbox' id='notify-when-quarantined-checkbox'/>
              </div>

              <div className={"inline flex items-center"}>
                <label htmlFor='quarantine-period-input'>Days before a quarantined message is deleted</label>
                <input type='number' id='quarantine-period-input'/>
              </div>

              <div className={"inline flex items-center"}>
                <label htmlFor='notify-when-quarantined-deleted-checkbox'>Notify me when a quarantined message is going to be deleted</label>
                <input type='checkbox' id='notify-when-quarantined-deleted-checkbox'/>
              </div>
            </div>
        </Section>
    );
};

export default Preferences;