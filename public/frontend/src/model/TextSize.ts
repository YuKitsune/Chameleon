
export enum TextSize {
	ExtraLarge,
	Large,
	NotLarge
}

export const textSizeToClassName = (size: TextSize): string => {
	switch (size) {
		case TextSize.ExtraLarge:
			return 'text-3xl';
		case TextSize.Large:
			return 'text-xl';
		case TextSize.NotLarge:
			return 'text-lg';
	}
}
