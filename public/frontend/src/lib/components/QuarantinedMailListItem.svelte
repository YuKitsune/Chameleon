<script lang="ts">
	import type QuarantinedMail from '$lib/model/quarantinedMail';
	export let quarantinedMail: QuarantinedMail;
</script>

<div class='flex justify-between rounded-md bg-gray-100 p-2 flex-row shadow-md cursor-pointer'>

	<div class='inline-flex grid grid-cols-2 gap-1 flex-grow content-center'>

		<!-- Top row -->
		<span class='inline-flex gap-1'>From: <strong><pre>{quarantinedMail.Sender}</pre></strong></span>
		<span class='text-right'>{quarantinedMail.DateReceived} days ago</span>

		<!-- Middle row -->
		<span class='inline-flex gap-1'>
			To: <strong class='inline-flex'>{quarantinedMail.Recipient.Name} (<pre>{quarantinedMail.Recipient.Address}</pre>)</strong>
		</span>

		<!-- Bottom row -->
		<span class='col-span-2 truncate'>{quarantinedMail.Subject}</span>
	</div>
</div>
