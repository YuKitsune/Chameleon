import type { AliasInfo } from '$lib/model/alias';

export default interface QuarantinedMail {
	Sender: string;
	Recipient: AliasInfo;
	Subject: string;
	DateReceived: number;
	VirusTotalRating: number;
	DomainTrustRating: number;
}
