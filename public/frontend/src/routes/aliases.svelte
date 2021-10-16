<script lang="ts">
	import AliasList from '$lib/AliasList.svelte';

	let dummyData = [
		{Alias: "john.doe@chameleon.io", WhitelistPattern: ".*@facebook\\.com"},
		{Alias: "jupiter_push@chameleon.io", WhitelistPattern: ".*@google\\.com"},
		{Alias: "andromeda-audio@chameleon.io", WhitelistPattern: ".*@somewhere-dodgy\\.com"}
	]
</script>
<svelte:head>
	<title>Aliases</title>
</svelte:head>

<div>
	<h1 class='text-xl'>Aliases</h1>

	<!-- Dummy data-->
	<AliasList aliases={dummyData} />
</div>