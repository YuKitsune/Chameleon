<script lang="ts">
	import AliasList from '$lib/components/AliasList.svelte';

	let dummyData = [
		{Name:"Facebook", Address: "john.doe@chameleon.io", WhitelistPattern: ".*@facebook\\.com", LastUsed: 5, EncryptionEnabled: true},
		{Name:"Google", Address: "jupiter_push@chameleon.io", WhitelistPattern: ".*@google\\.com", LastUsed: 3, EncryptionEnabled: true},
		{Name:"Somewhere Dodgy", Address: "andromeda-audio@chameleon.io", WhitelistPattern: ".*@somewhere-dodgy\\.com", LastUsed: 2, EncryptionEnabled: false}
	]
</script>
<svelte:head>
	<title>Aliases</title>
</svelte:head>

<div>
	<h1 class='text-2xl'>Aliases</h1>
	<AliasList aliases={dummyData} />
</div>