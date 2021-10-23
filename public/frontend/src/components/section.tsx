import React, { PropsWithChildren } from 'react';
import { TextSize, textSizeToClassName } from '../model/TextSize';
import PropsWithClassName from './PropsWithClassName';

type SectionProps = PropsWithChildren<{
    header: string;
    headerSize?: TextSize;
 }> & PropsWithClassName;

const Section = (props: SectionProps) => {

    let headerSize = props.headerSize ?? TextSize.ExtraLarge;

    return (
        <div className={`flex flex-col gap-1 pt-2 ${props.className}`}>
            <h1 className={`${textSizeToClassName(headerSize)} border-b border-gray-300`}>
                {props.header}
            </h1>
            <div className='p-1'>
                {props.children}
            </div>
        </div>
    );
};

export default Section;