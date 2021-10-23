import type { AliasInfo } from './alias';

export interface QuarantinedMail {
	Sender: string;
	Recipient: AliasInfo;
	Subject: string;
	DateReceived: number;
	VirusTotalRating: number;
	DomainTrustRating: number;
}
