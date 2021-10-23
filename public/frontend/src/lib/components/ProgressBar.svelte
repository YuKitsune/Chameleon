<script lang="ts">
	const maxPercent = 100;

	export let value = 0;
	export let max = 100;
	export let className: string;

	let percentage = 0;

	const calculatePercentage = (): number => {
		if (value <= 0) return 0;

		let val = (value / max) * 100;
		return Math.min(val, maxPercent)
	}

	$: {
		percentage = calculatePercentage();
	}
</script>

<div class="w-full bg-gray-200 rounded-md {className}">
	<div class="bg-green-500 h-full rounded-md" style="width: {percentage}%"></div>
</div>