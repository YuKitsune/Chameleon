<script lang="ts">
	export let isOpen = false;
	const toggleDropdown = () => isOpen = !isOpen;

	const hideDropdown = () => {
		isOpen = false;
	}

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

		hideDropdown();
	}
</script>

<svelte:body on:click={onClickOutside} />

<div class="relative" bind:this={dropdownElement}>
	<div on:click={toggleDropdown}>
		<slot name='button' />
	</div>
	{#if isOpen}
		<div class='absolute right-0 mt-2 py-2 w-40 bg-white rounded-md shadow-md' on:click={hideDropdown}>
			<slot name='items' />
		</div>
	{/if}
</div>