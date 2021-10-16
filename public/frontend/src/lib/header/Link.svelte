<script>
	import { page } from '$app/stores';
	export let path;
</script>

<div class='mx-1 link' class:active={$page.path === path}>
	<a href={path}>
		<slot/>
	</a>
</div>

<style>
	.link {
		@apply px-2 py-1 rounded-md cursor-pointer;
	}

  .link:hover {
		@apply bg-green-600;
  }

  .link.active {
		@apply bg-green-600;
  }

</style>