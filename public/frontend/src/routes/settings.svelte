<script>
	import Section from '$lib/Section.svelte';
</script>

<div class='flex flex-row grid grid-cols-2 gap-4'>

	<Section header='Mail Forwarding'>
		<!-- Retry interval -->
		<!-- Number of retries before giving up -->
	</Section>

	<Section header='Quarantined Mail'>
		<!-- Days to keep before deleting -->
	</Section>

	<Section header='Danger Zone'>
		<!-- Delete all quarantined mail -->
		<!-- Delete all aliases -->
	</Section>

</div>
