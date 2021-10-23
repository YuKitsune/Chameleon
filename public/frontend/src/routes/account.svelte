<script>
	import TextInput from '$lib/components/TextInput.svelte';
	import {TextSize} from '$lib/model/TextSize';
	import Section from '$lib/components/Section.svelte';
	import Button from '$lib/components/Button.svelte';
	import PayPalIcon from '$lib/components/icons/PayPalIcon.svelte';
	import ProgressBar from '$lib/components/ProgressBar.svelte';
</script>

<div class='two-column-grid gap-4'>

	<!-- Details -->
	<Section header='My Details'>
		<TextInput label='Primary Email'/>
		<TextInput label='Secondary Email'/>
	</Section>

	<!-- Billing -->
	<Section header='Billing'>
		<div class='two-column-grid'>

			<Section header='Plan' headerSize={TextSize.Large}>
				<div class='two-column-grid'>

					<!-- Plan title -->
					<span><strong>Premium +</strong></span>

					<!-- Change plan button -->
					<div class='inline-flex justify-end'>
						<Button className='bg-gray-200 hover:bg-gray-300'>Change plan</Button>
					</div>

					<!-- Next billing date -->
					<span class='col-span-2'>Next billing date: 1/12/2022</span>

					<div>
						<label for='auto-renew-checkbox'>Auto-renew</label>
						<input type='checkbox' id='auto-renew-checkbox'/>
					</div>
				</div>
			</Section>

			<Section header='Payment Method' headerSize={TextSize.Large}>
				<div class='one-column-grid'>

					<!-- Selected payment method -->
					<div class='flex inline-flex align-middle'>
						<PayPalIcon className='w-16 mt-1 mr-2' /> <strong>****doe@mymail.com</strong>
					</div>

					<!-- Change Payment Method button -->
					<!-- Todo: Let's not use a button here... It's kinda wide -->
					<div class='inline-flex justify-end'>
						<Button className='bg-gray-200 hover:bg-gray-300'>Change payment method</Button>
					</div>
				</div>
			</Section>

			<!-- Usage -->
			<Section header='Usage' headerSize={TextSize.Large} className='col-span-2'>
				<div class='two-column-grid'>
					<div class='flex flex-col'>
						<strong>Cached Mail</strong>
						<ProgressBar value={150} max={1024} className='h-3 m-1' />
					</div>
					<div class='flex flex-col'>
						<strong>Quarantined Mail</strong>
						<ProgressBar value={256} max={1024} className='h-3 m-1' />
					</div>
				</div>
			</Section>

		</div>
	</Section>

	<!-- Security -->
	<Section header='Security' className='col-span-2'>
		<div class='two-column-grid'>

			<Section header='Password' headerSize={TextSize.Large}>
				<TextInput label='Password' isSensitive={true}/>
				<TextInput label='Confirm Password' isSensitive={true}/>
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
	<Section header='Danger Zone' className='col-span-2'>

		<div class='inline-flex'>
			<!-- Delete account -->
			<Button className='bg-red-500 hover:bg-red-600'>Delete account</Button>
		</div>

	</Section>

</div>