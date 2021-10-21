<script lang="ts">
	import { page } from '$app/stores';
	export let path: string | undefined = undefined;

	// Reeeeeeee why can't i just put a class on the component???
	export let className: string | undefined = undefined;
</script>

<div class='link {className}' class:active={$page.path === path}>
	<a href={path}>
		<slot/>
	</a>
</div>

<style>
	.link {
		@apply px-2 py-1 rounded-md cursor-pointer;
	}

  .link:hover {
		@apply bg-green-600 text-white;
  }

  .link.active {
		@apply bg-green-600;
  }
</style>