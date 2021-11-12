import React from 'react';
import Section from '../section';
import { TextSize } from '../../model/TextSize';

const Preferences = () => {

    return (
      <Section header='Preferences'>
        <div className='flex flex-row grid grid-cols-2 gap-4'>

            <Section header='Mail Forwarding' headerSize={TextSize.Large}>
                {/* Retry interval */}
                {/* Number of retries before giving up */}
            </Section>

            <Section header='Quarantined Mail' headerSize={TextSize.Large}>
                {/* Days to keep before deleting */}
            </Section>

            <Section header='Danger Zone' headerSize={TextSize.Large}>
                {/* Delete all quarantined mail */}
                {/* Delete all aliases */}
            </Section>

        </div>
      </Section>
    );
};

export default Preferences;