<script>
	import UserCircle from '$lib/icons/UserCircle.svelte';
	import NotificationIcon from '$lib/icons/NotificationIcon.svelte';
	import Link from '$lib/header/Link.svelte';
	import Dropdown from '$lib/Dropdown.svelte';
	import DropdownItem from '$lib/DropdownItem.svelte';
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
			<Link className='mx-2' path='/aliases'>Aliases</Link>
			<Link className='mx-1' path='/quarantined'>Quarantined</Link>
		</div>
	</div>

	<!-- Right -->
	<!-- Todo: These should be dropdowns -->
	<div class="inline-flex items-center">
		<Link className='mx-1' path='/notifications'><NotificationIcon /></Link>

		<Dropdown>
			<span slot='button'>
				<Link className='mx-1'><UserCircle /></Link>
			</span>
			<span slot='items'>
				<DropdownItem>
					<a href='/account'>My Account</a>
				</DropdownItem>
				<DropdownItem>
					<a href='/settings'>Settings</a>
				</DropdownItem>
				<DropdownItem>
					<a href='/logout'>Log Out</a>
				</DropdownItem>
			</span>
		</Dropdown>
	</div>
</div>