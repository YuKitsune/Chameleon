
<script lang="ts">
	import AliasRow from '$lib/AliasRow.svelte';
	import type Alias from '$lib/alias';

	export let aliases: Alias[] = [];
</script>

<div class='flex flex-col gap-1'>
	{#each aliases as alias}
		<AliasRow alias={alias} />
	{/each}
</div>