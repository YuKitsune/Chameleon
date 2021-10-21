<script lang="ts">
	export let label: string;
	export let placeholder: string;
	export let name: string | undefined;
	export let className: string | undefined;
</script>

<div class='py-1 {className}'>
	<label class="block text-gray-700 text-sm font-bold mb-2" for={name}>
		{label}
	</label>
	<input class="shadow appearance-none border rounded-sm w-full py-1 px-2 text-gray-700 leading-tight focus:outline-none focus:shadow-outline focus:border-green-500" type="text" placeholder={placeholder} id={name}>
</div>