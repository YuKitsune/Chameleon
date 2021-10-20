<script>
	import TextInput from '$lib/TextInput.svelte';
	import {TextSize} from '$lib/TextSize';
	import Section from '$lib/Section.svelte';
</script>

<div class='flex flex-row grid grid-cols-2 gap-4'>

	<!-- Details -->
	<Section header='My Details'>
		<TextInput label='Primary Email'/>
		<TextInput label='Secondary Email'/>
	</Section>

	<!-- Billing -->
	<Section header='Billing'>
		<div class='flex flex-row grid grid-cols-2 gap-2'>

			<Section header='Tiers' headerSize={TextSize.Large}>
				🚧 Todo 🚧
			</Section>

			<Section header='Payment Method' headerSize={TextSize.Large}>
				🚧 Todo 🚧
			</Section>

			<Section header='Usage' headerSize={TextSize.Large} className='col-span-2'>
				🚧 Todo 🚧
			</Section>

		</div>
	</Section>

	<!-- Security -->
	<Section header='Security'>
		<div class='flex flex-row grid grid-cols-2 gap-2'>

			<Section header='Password' headerSize={TextSize.Large}>
				🚧 Todo 🚧
			</Section>

			<Section header='Two Factor Authentication' headerSize={TextSize.Large}>
				🚧 Todo 🚧
			</Section>

			<Section header='Sessions' headerSize={TextSize.Large}>
				🚧 Todo 🚧
			</Section>

			<Section header='Events' headerSize={TextSize.Large}>
				🚧 Todo 🚧
			</Section>

		</div>
	</Section>

	<!-- Danger Zone -->
	<Section header='Danger Zone'>
		🚧 Todo 🚧
		<!-- Delete account -->
	</Section>

</div>
