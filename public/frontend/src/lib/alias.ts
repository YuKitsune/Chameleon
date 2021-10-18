
export default interface Alias {
	Name: string;
	Address: string;
	WhitelistPattern: string;
	IsActive: boolean;
	EncryptionEnabled: boolean;
	LastUsed: number;
}

export type AliasTemplate = "Template";