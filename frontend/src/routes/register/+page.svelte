<script>
	import { goto } from '$app/navigation';
	let name = '';
	let email = '';
	let password = '';
	let confirmPassword = '';
	let errorMessage = '';
	let successMessage = '';
	const passwordRegex = /^(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{6,}$/; // Password harus mengandung huruf kapital, angka, dan karakter khusus.
	let showPassword = false;
	let showConfirmPassword = false;

	const handleRegister = async () => {
		if (!name || !email || !password || !confirmPassword) {
			errorMessage = 'Semua kolom wajib diisi';
			successMessage = '';
		} else if (password !== confirmPassword) {
			errorMessage = 'Password dan konfirmasi password tidak cocok';
			successMessage = '';
		} else if (!passwordRegex.test(password)) {
			errorMessage =
				'Password harus memiliki minimal 6 karakter, mengandung huruf kapital, angka, dan karakter khusus.';
			successMessage = '';
		} else {
			try {
				const response = await fetch('http://localhost:8080/auth/register', {
					method: 'POST',
					headers: {
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({
						name,
						email,
						password,
						role: 'user' // Kirim role, jika ingin mengaturnya secara eksplisit
					})
				});

				const data = await response.json();

				if (response.ok) {
					successMessage = 'Pendaftaran berhasil! Silakan login.';
					errorMessage = '';
					// Redirect ke halaman login setelah berhasil daftar
					goto('/login');
				} else {
					errorMessage = data.error || 'Pendaftaran gagal, coba lagi';
					successMessage = '';
				}
			} catch (err) {
				errorMessage = 'Terjadi kesalahan, coba lagi';
				successMessage = '';
			}
		}
	};

	const togglePassword = () => {
		showPassword = !showPassword;
	};

	const toggleConfirmPassword = () => {
		showConfirmPassword = !showConfirmPassword;
	};
</script>

<div class="flex min-h-screen items-center justify-center bg-gray-100">
	<div class="w-full max-w-sm rounded-lg bg-white p-8 shadow-lg">
		<h2 class="mb-4 text-center text-2xl font-semibold">Daftar</h2>

		{#if errorMessage}
			<div class="mb-4 rounded bg-red-500 p-2 text-center text-white">
				{errorMessage}
			</div>
		{/if}

		{#if successMessage}
			<div class="mb-4 rounded bg-green-500 p-2 text-center text-white">
				{successMessage}
			</div>
		{/if}

		<form on:submit|preventDefault={handleRegister}>
			<div class="mb-4">
				<label for="name" class="block text-sm font-medium text-gray-700">Nama</label>
				<input
					id="name"
					type="text"
					bind:value={name}
					class="mt-1 w-full rounded-md border border-gray-300 p-2"
					placeholder="Nama"
				/>
			</div>

			<div class="mb-4">
				<label for="email" class="block text-sm font-medium text-gray-700">Email</label>
				<input
					id="email"
					type="email"
					bind:value={email}
					class="mt-1 w-full rounded-md border border-gray-300 p-2"
					placeholder="Email"
				/>
			</div>

			<div class="relative mb-4">
				<label for="password" class="block text-sm font-medium text-gray-700">Password</label>
				<input
					id="password"
					type={showPassword ? 'text' : 'password'}
					bind:value={password}
					class="mt-1 w-full rounded-md border border-gray-300 p-2"
					placeholder="Password"
				/>
				<button
					type="button"
					class="absolute inset-y-0 right-0 flex items-center pr-3"
					on:click={togglePassword}
				>
					{#if showPassword}
						<span>👁️</span>
					{:else}
						<span>👁️‍🗨️</span>
					{/if}
				</button>
			</div>

			<div class="relative mb-4">
				<label for="confirmPassword" class="block text-sm font-medium text-gray-700">
					Konfirmasi Password
				</label>
				<input
					id="confirmPassword"
					type={showConfirmPassword ? 'text' : 'password'}
					bind:value={confirmPassword}
					class="mt-1 w-full rounded-md border border-gray-300 p-2"
					placeholder="Konfirmasi Password"
				/>
				<button
					type="button"
					class="absolute inset-y-0 right-0 flex items-center pr-3"
					on:click={toggleConfirmPassword}
				>
					{#if showConfirmPassword}
						<span>👁️</span>
					{:else}
						<span>👁️‍🗨️</span>
					{/if}
				</button>
			</div>

			<button
				type="submit"
				class="w-full rounded-md bg-blue-500 p-2 text-white hover:bg-blue-600 focus:outline-none"
			>
				Daftar
			</button>
		</form>

		<p class="mt-4 text-center text-sm">
			Sudah punya akun?
			<a href="/login" class="text-blue-500 hover:underline">Masuk</a>
		</p>
	</div>
</div>
