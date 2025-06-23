<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	type Employee = {
		id: number;
		name: string;
		npwp: string;
		address: string;
	};

	let employees: Employee[] = [];

	onMount(async () => {
		const token = localStorage.getItem('authToken');
		const expiresAt = parseInt(localStorage.getItem('tokenExpiresAt') || '0', 10);
		const now = Date.now();

		if (!token || now >= expiresAt) {
			alert('Sesi Anda telah berakhir. Silakan login kembali.');
			localStorage.clear();
			goto('/login');
			return;
		}

		try {
			const res = await fetch('http://localhost:8080/index_employee', {
				headers: {
					Authorization: `Bearer ${token}`
				}
			});
			const data = await res.json();

			if (data && data.employees) {
				employees = data.employees;
			} else {
				console.error('Data employees tidak ditemukan');
			}
		} catch (err) {
			console.error('Gagal mengambil data:', err);
		}
	});

	const deleteEmployee = async (id: number) => {
		const confirmation = confirm('Apakah Anda yakin ingin menghapus data ini?');
		if (confirmation) {
			const token = localStorage.getItem('authToken');
			if (!token) {
				alert('Token tidak ditemukan. Silakan login ulang.');
				goto('/login');
				return;
			}

			const res = await fetch(`http://localhost:8080/delete_employee/${id}`, {
				method: 'DELETE',
				headers: {
					Authorization: `Bearer ${token}`
				}
			});

			if (res.ok) {
				employees = employees.filter((emp) => emp.id !== id);
			} else {
				console.error('Gagal menghapus data');
			}
		}
	};
</script>

<div class="container mx-auto p-6">
	<div class="mb-4 flex items-center justify-between">
		<h3 class="text-2xl font-semibold">List Employee</h3>
		<a
			href="/create_employee"
			class="rounded bg-blue-500 px-4 py-2 text-white transition hover:bg-blue-600"
		>
			Tambah Data
		</a>
	</div>

	<table class="min-w-full table-auto rounded-lg border border-gray-300 bg-white shadow-md">
		<thead>
			<tr class="bg-gray-100">
				<th class="px-4 py-2 text-left text-sm font-medium text-gray-700">Nama</th>
				<th class="px-4 py-2 text-left text-sm font-medium text-gray-700">NPWP</th>
				<th class="px-4 py-2 text-left text-sm font-medium text-gray-700">Alamat</th>
				<th class="px-4 py-2 text-left text-sm font-medium text-gray-700">Action</th>
			</tr>
		</thead>
		<tbody>
			{#each employees as emp}
				<tr class="border-b border-gray-200">
					<td class="px-4 py-2 text-sm text-gray-900">{emp.name}</td>
					<td class="px-4 py-2 text-sm text-gray-900">{emp.npwp}</td>
					<td class="px-4 py-2 text-sm text-gray-900">{emp.address}</td>
					<td class="px-4 py-2">
						<a
							href={`/update_employee/${emp.id}`}
							class="rounded bg-yellow-500 px-3 py-1 text-white transition hover:bg-yellow-600"
							>Edit</a
						>
						<button
							on:click={() => deleteEmployee(emp.id)}
							class="ml-2 rounded bg-red-500 px-3 py-1 text-white transition hover:bg-red-600"
							>Hapus</button
						>
					</td>
				</tr>
			{/each}
		</tbody>
	</table>
</div>
