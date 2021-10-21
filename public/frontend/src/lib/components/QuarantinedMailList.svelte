<script lang='ts'>
	import QuarantinedMail from '$lib/model/quarantinedMail';
	import QuarantinedMailListItem from '$lib/components/QuarantinedMailListItem.svelte';

	export let quarantinedMail: QuarantinedMail[] = [];
</script>

<!-- Todo: This and the ListItem need some re-work -->
<div class='flex flex-col grid grid-cols-2 gap-4'>
	{#each quarantinedMail as qMail}
		<QuarantinedMailListItem quarantinedMail={qMail} />
	{/each}
</div>
