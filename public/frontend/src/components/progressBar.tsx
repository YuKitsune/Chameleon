import React, { useState, useCallback, useEffect } from 'react';
import PropsWithClassName from './PropsWithClassName';

type ProgressBarProps = PropsWithClassName & {
    value: number;
    max: number;
};

const ProgressBar = (props: ProgressBarProps) => {

    const maxPercent = 100;

    const [percentage, setPercentage] = useState(0);

    const calculatePercentage = useCallback(() => {
		if (props.value <= 0) return 0;

		let value = (props.value / props.max) * 100;
		return Math.min(value, maxPercent)
    }, [props])

    useEffect(() => {
        setPercentage(calculatePercentage())
    }, [props.value, props.max])

    return (
        <div className={`w-full bg-gray-200 rounded-md ${props.className}`}>
            <div className="bg-green-500 h-full rounded-md" style={{ width: percentage }}></div>
        </div>
    );
};

export default ProgressBar;