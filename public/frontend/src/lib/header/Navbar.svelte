<script>
	import UserCircle from '$lib/icons/UserCircle.svelte';
	import NotificationIcon from '$lib/icons/NotificationIcon.svelte';
	import Link from '$lib/header/Link.svelte';
</script>

<div class="items-center justify-between flex bg-green-500 text-white bg-opacity-90 px-24 py-2 shadow-2xl">

	<!-- Left -->
	<div class="inline-flex items-center">

		<!-- Todo: Logo -->
		<div class="text-2xl text-white font-semibold inline-flex">
			<span>Chameleon</span>
		</div>

		<!-- Links -->
		<div class='ml-5 flex'>
			<Link path='/'>Home</Link>
			<Link path='/aliases'>Aliases</Link>
			<Link path='/quarantined'>Quarantined</Link>
		</div>
	</div>

	<!-- Right -->
	<!-- Todo: These should be dropdowns -->
	<div class="inline-flex items-center">
		<Link path='/notifications'><NotificationIcon /></Link>
		<Link path='/account'><UserCircle /></Link>
	</div>

</div>