<script>
	import { goto } from '$app/navigation';

	let email = '';
	let password = '';
	let showPassword = false;
	let errorMessage = '';
	let successMessage = '';

	const handleLogin = async () => {
		if (!email || !password) {
			errorMessage = 'Email dan password wajib diisi';
			successMessage = '';
			return;
		}

		try {
			const response = await fetch('http://localhost:8080/auth/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ email, password })
			});

			const data = await response.json();

			if (response.ok) {
				// Simpan token
				localStorage.setItem('authToken', data.token);
				localStorage.setItem('role', data.user.role);

				// Ambil exp dari token
				const base64Url = data.token.split('.')[1];
				const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
				const decoded = JSON.parse(atob(base64));
				const expiresAt = decoded.exp * 1000;
				localStorage.setItem('tokenExpiresAt', expiresAt.toString());

				// Logout otomatis saat token expired
				const remainingTime = expiresAt - Date.now();
				setTimeout(() => {
					alert('Sesi telah berakhir, silakan login kembali.');
					localStorage.clear();
					goto('/login');
				}, remainingTime);

				// Arahkan ke halaman sesuai role
				successMessage = `Login berhasil sebagai ${data.user.role}`;
				errorMessage = '';

				goto('/index_employee');
			} else {
				errorMessage = data.error || 'Login gagal, coba lagi';
				successMessage = '';
			}
		} catch (err) {
			errorMessage = 'Terjadi kesalahan, coba lagi';
			successMessage = '';
		}
	};

	const togglePasswordVisibility = () => {
		showPassword = !showPassword;
	};
</script>

<div class="flex min-h-screen items-center justify-center bg-gray-100">
	<div class="w-full max-w-sm rounded-lg bg-white p-8 shadow-lg">
		<h2 class="mb-4 text-center text-2xl font-semibold">Login</h2>

		{#if errorMessage}
			<div class="mb-4 rounded bg-red-500 p-2 text-center text-white">{errorMessage}</div>
		{/if}

		<form on:submit|preventDefault={handleLogin}>
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

			<div class="mb-4">
				<label for="password" class="block text-sm font-medium text-gray-700">Password</label>
				<div class="relative">
					<input
						id="password"
						type={showPassword ? 'text' : 'password'}
						bind:value={password}
						class="mt-1 w-full rounded-md border border-gray-300 p-2"
						placeholder="Password"
					/>
					<button
						type="button"
						class="absolute right-2 top-1/2 -translate-y-1/2 transform text-gray-500"
						on:click={togglePasswordVisibility}
					>
						{#if showPassword}
							🙈
						{:else}
							👁️
						{/if}
					</button>
				</div>
			</div>

			<button type="submit" class="w-full rounded-md bg-blue-500 p-2 text-white hover:bg-blue-600">
				Login
			</button>
		</form>

		<p class="mt-4 text-center text-sm">
			Belum punya akun?
			<a href="/register" class="text-blue-500 hover:underline">Daftar</a>
		</p>
	</div>
</div>
