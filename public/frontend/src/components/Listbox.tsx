import React, { useEffect, useState } from 'react';
import PropsWithClassName from './PropsWithClassName';
import {Listbox as HeadlessListbox} from '@headlessui/react';

interface ListboxItem {
	Label: string;
	Value: any;
	Disabled: boolean;
}

type OnSelectionChangedCallback = (item: ListboxItem) => void;

type ListboxProps = PropsWithClassName & {
	items: ListboxItem[];
	onSelectionChanged?: OnSelectionChangedCallback | undefined;
};

const Listbox = (props: ListboxProps) => {
	const [selectedItem, setSelectedItem] = useState<ListboxItem>(props.items[0]);

	useEffect(() => {
		if (!props.onSelectionChanged)
			return;

		props.onSelectionChanged(selectedItem);
	}, [selectedItem]);

	return (
		<div className={props.className}>
			<HeadlessListbox value={selectedItem} onChange={setSelectedItem}>
				<HeadlessListbox.Button>{selectedItem.Label}</HeadlessListbox.Button>
				<HeadlessListbox.Options>
					{props.items.map((item) => (
						<HeadlessListbox.Option
							key={item.Value}
							value={item.Value}
							disabled={item.Disabled}
						>
							{item.Label}
						</HeadlessListbox.Option>
					))}
				</HeadlessListbox.Options>
			</HeadlessListbox>
		</div>
	);
}

export default Listbox;