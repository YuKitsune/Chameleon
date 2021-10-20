<script lang="ts">
	export let isOpen = false;

	const toggleIsOpen = () => isOpen = !isOpen;
</script>

<div class="relative">
	<div on:click={toggleIsOpen}>
		<slot name='button' />
	</div>
	{#if isOpen}
		<div class='absolute right-0 mt-2 py-2 w-40 bg-white rounded-md shadow-md'>
			<slot name='items' />
		</div>
	{/if}
</div>