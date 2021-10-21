
<script lang="ts">
	import AliasListItem from '$lib/components/AliasListItem.svelte';
	import type Alias from '$lib/model/alias';

	export let aliases: Alias[] = [];
	export let allowAdd = true;
</script>

<div class='flex flex-row grid grid-cols-2 gap-4'>
	{#each aliases as alias}
		<AliasListItem alias={alias} />
	{/each}

	{#if allowAdd}
		<AliasListItem isTemplate={true} />
	{/if}
</div>