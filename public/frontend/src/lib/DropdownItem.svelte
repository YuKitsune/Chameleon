<div class='dropdown-item'>
	<slot/>
</div>

<style>
    .dropdown-item {
        @apply px-2 py-1 cursor-pointer text-black;
    }

    .dropdown-item:hover {
        @apply bg-green-600 text-white;
    }
</style>