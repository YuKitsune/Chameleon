<script lang="ts">
	import Alias from "./alias"
	import RemoveIcon from '$lib/icons/RemoveIcon.svelte';
	export let alias: Alias;
</script>


<div class='flex justify-between rounded-md bg-gray-100 p-2 flex-row'>

	<!-- Left: Info -->
	<div class='inline-flex'>
		<span>{alias.Alias}</span>
		<span>{alias.WhitelistPattern}</span>
	</div>

	<!-- Right: Actions -->
	<div class='inline-flex'>
		<span class='rounded-md bg-red-500 text-white'>
			<RemoveIcon />
		</span>
	</div>
</div>