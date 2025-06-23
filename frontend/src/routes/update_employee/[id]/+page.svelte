<script lang="ts">
	import { goto } from '$app/navigation';

	import { onMount } from 'svelte';
	import { getIdEmployee, updateEmployee } from '../../../lib/api_employee';
	import { validatorEmployee } from '../../../lib/validator_employee';

	export let data: { id: string };
	let { id } = data;

	let name = '';
	let npwp = '';
	let address = '';

	let errorName = '';
	let errorNPWP = '';
	let errorAddress = '';

	// Ambil data karyawan berdasarkan ID
	onMount(async () => {
		const employeeData = await getIdEmployee(id);
		name = employeeData.name;
		npwp = employeeData.npwp;
		address = employeeData.address;
	});

	// Fungsi untuk mengirim data yang sudah diperbarui
	const handleSubmit = async (e: Event) => {
		e.preventDefault();

		const isValid = await validatorEmployee(name, npwp, address, (errors) => {
			errorName = errors.name;
			errorNPWP = errors.npwp;
			errorAddress = errors.address;
		});

		if (!isValid) {
			return; // Jika validasi gagal, hentikan eksekusi
		}

		await updateEmployee(id, name, npwp, address);
		alert('Data berhasil diupdate');
		goto('/index_employee'); // Redirect ke halaman yang diinginkan
	};
</script>

<div class="container mx-auto max-w-xl p-6">
	<h3 class="mb-4 text-2xl font-semibold">Edit Karyawan</h3>

	<form on:submit|preventDefault={handleSubmit} class="space-y-4">
		<div>
			<label class="block text-sm font-medium text-gray-700" for="name">Nama Karyawan</label>
			<input
				id="name"
				type="text"
				bind:value={name}
				class="mt-1 w-full rounded-md border px-4 py-2 shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
			/>
			{#if errorName}
				<div class="text-sm text-red-600">{errorName}</div>
			{/if}
		</div>

		<div>
			<label class="block text-sm font-medium text-gray-700" for="npwp">NPWP</label>
			<input
				id="npwp"
				type="text"
				bind:value={npwp}
				class="mt-1 w-full rounded-md border px-4 py-2 shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
			/>
			{#if errorNPWP}
				<div class="text-sm text-red-600">{errorNPWP}</div>
			{/if}
		</div>

		<div>
			<label class="block text-sm font-medium text-gray-700" for="address">Alamat</label>
			<textarea
				id="address"
				bind:value={address}
				class="mt-1 w-full rounded-md border px-4 py-2 shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
			></textarea>
			{#if errorAddress}
				<div class="text-sm text-red-600">{errorAddress}</div>
			{/if}
		</div>

		<button
			type="submit"
			class="w-full rounded-md bg-blue-500 px-4 py-2 text-white transition hover:bg-blue-600"
		>
			Simpan Perubahan
		</button>
	</form>
</div>
