<script lang="ts">
	import Alias from './alias';
	import AddIcon from '$lib/icons/AddIcon.svelte';
	import LockIcon from '$lib/icons/LockIcon.svelte';
	export let alias: Alias | undefined = undefined;
	export let isTemplate = false;
</script>

<div class='flex justify-between rounded-md bg-gray-100 p-2 flex-row shadow-md cursor-pointer'>

	<div class='inline-flex grid grid-cols-2 gap-1 flex-grow content-center'>
		{#if isTemplate}
			<span class='text-xl row-span-2 text-right'>Add</span>
			<AddIcon class='row-span-2'/>
		{/if}

		{#if alias}
			<!-- Top row -->
			<span class='inline-flex gap-2'>
				<strong>{alias.Name}</strong>
				{#if alias.EncryptionEnabled}
					<!-- Todo: This doesn't quite line up -->
					<LockIcon className='stroke-current text-green-500 h-5 w-5'/>
				{/if}
			</span>
			<span>Last used: {alias.LastUsed} days ago</span>

			<!-- Bottom row -->
			<span><pre>{alias.Address}</pre></span>
		{/if}
	</div>
</div>