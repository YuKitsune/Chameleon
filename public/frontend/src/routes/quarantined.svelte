<script lang="ts">
	import QuarantinedMailList from '$lib/QuarantinedMailList.svelte';

	let dummyData = [
		{Sender:"danny.devito@macaroni.cheese", Recipient: {Name:"Facebook", Address: "john.doe@chameleon.io"}, Subject: "Smell my nuts", DateReceived: 5},
		{Sender:"sammy.von.soomel@nintento.info", Recipient: {Name:"Facebook", Address: "john.doe@chameleon.io"}, Subject: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.", DateReceived: 2},
	]
</script>

<svelte:head>
	<title>Quarantined</title>
</svelte:head>

<div>
	<h1 class='text-2xl'>Quarantined</h1>
	<QuarantinedMailList quarantinedMail={dummyData} />
</div>