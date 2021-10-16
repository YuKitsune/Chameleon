<script lang="ts">
	import Navbar from '$lib/header/Navbar.svelte';
	import '../tailwind.css'
</script>

<div class='flex flex-col h-screen overflow-hidden bg-gray-100'>
	<Navbar class='w-full'/>

	<div class='content'>
		<main class='child'>
			<slot/>
		</main>

		<footer class='footer'>
			<p>visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to learn SvelteKit</p>
		</footer>

	</div>
</div>


<style>
	.content {
			@apply flex flex-col flex-grow items-center justify-between px-24 p-4 overflow-y-scroll;
	}

	.child {
			@apply flex-grow w-full rounded-md bg-white shadow-2xl p-4;
	}

	.footer {
			@apply flex flex-col flex-grow-0 items-center justify-center pt-4;
	}

	footer a {
		font-weight: bold;
	}
</style>
