<script lang="ts">
	import { validatorEmployee } from '../../lib/validator_employee';
	import { createEmployee } from '../../lib/api_employee';
	import { goto } from '$app/navigation';

	let name = '';
	let npwp = '';
	let address = '';

	let errorName = '';
	let errorNPWP = '';
	let errorAddress = '';

	const validator = 
	async (event: Event) => {
		event.preventDefault();
		// Validasi input sebelum mengirim data
		const isValid = await validatorEmployee(name, npwp, address, (errors) => {
			errorName = errors.name;
			errorNPWP = errors.npwp;
			errorAddress = errors.address;
		});

		if (!isValid) {
			return; // Jika validasi gagal, hentikan eksekusi
		}

		// Jika validasi berhasil, kirim data ke API
		const response = await createEmployee(name, npwp, address);
		if (response.ok) {
			goto('/index_employee'); // Redirect ke halaman yang diinginkan
		}
	};
</script>

<div class="container mx-auto max-w-3xl p-6">
	<form on:submit={validator}>
		<table class="w-full table-auto border-separate border-spacing-y-4">
			<tbody>
				<tr>
					<td class="w-1/4 align-top">
						<label for="name" class="mt-2 block text-sm font-medium text-gray-700"
							>Nama Karyawan</label
						>
					</td>
					<td>
						<input
							type="text"
							id="name"
							bind:value={name}
							class="w-full rounded-md border border-gray-300 px-4 py-2 shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
						/>
						{#if errorName}
							<div class="mt-1 text-sm text-red-600">{errorName}</div>
						{/if}
					</td>
				</tr>

				<tr>
					<td class="align-top">
						<label for="npwp" class="mt-2 block text-sm font-medium text-gray-700">NPWP</label>
					</td>
					<td>
						<input
							type="text"
							id="npwp"
							bind:value={npwp}
							class="w-full rounded-md border border-gray-300 px-4 py-2 shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
						/>
						{#if errorNPWP}
							<div class="mt-1 text-sm text-red-600">{errorNPWP}</div>
						{/if}
					</td>
				</tr>

				<tr>
					<td class="align-top">
						<label for="address" class="mt-2 block text-sm font-medium text-gray-700">Alamat</label>
					</td>
					<td>
						<textarea
							id="address"
							bind:value={address}
							class="w-full rounded-md border border-gray-300 px-4 py-2 shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
						></textarea>
						{#if errorAddress}
							<div class="mt-1 text-sm text-red-600">{errorAddress}</div>
						{/if}
					</td>
				</tr>

				<tr>
					<td></td>
					<td>
						<button
							type="submit"
							class="mt-4 w-full rounded-md bg-blue-500 px-4 py-2 text-white transition hover:bg-blue-600"
						>
							Tambah
						</button>
					</td>
				</tr>
			</tbody>
		</table>
	</form>
</div>
