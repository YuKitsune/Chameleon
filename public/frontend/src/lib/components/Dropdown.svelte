<script lang="ts">
	export let isOpen = false;
	const toggleIsOpen = () => isOpen = !isOpen;

	let dropdownElement;
	function isInDropdown(target) {
		let parent = target;
		while (parent) {
			if (parent === dropdownElement)
				return true;

			parent = parent.parentNode;
		}

		return false;
	}

	const onClickOutside = (event) => {
		if (isInDropdown(event.target))
			return;

		toggleIsOpen();
	}
</script>

<svelte:body on:click={onClickOutside} />

<div class="relative" bind:this={dropdownElement}>
	<div on:click={toggleIsOpen}>
		<slot name='button' />
	</div>
	{#if isOpen}
		<div class='absolute right-0 mt-2 py-2 w-40 bg-white rounded-md shadow-md'>
			<slot name='items' />
		</div>
	{/if}
</div>