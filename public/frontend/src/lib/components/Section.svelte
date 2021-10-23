<script lang='ts'>
	import { TextSize, textSizeToClassName } from '$lib/model/TextSize';

	export let header: string;
	export let headerSize: TextSize = TextSize.ExtraLarge;
	export let className: string;
</script>

<div class='flex flex-col gap-1 pt-2 {className}'>
	<h1 class='{textSizeToClassName(headerSize)} border-b border-gray-300'>{header}</h1>
	<div class='p-1'>
		<slot />
	</div>
</div>