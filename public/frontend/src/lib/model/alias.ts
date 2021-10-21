
export interface AliasInfo {
	Name: string;
	Address: string;
}

export default interface Alias extends AliasInfo {
	WhitelistPattern: string;
	IsActive: boolean;
	EncryptionEnabled: boolean;
	LastUsed: number;
}
