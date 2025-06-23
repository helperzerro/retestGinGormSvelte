// src/lib/api.ts

export const createEmployee = async (name: string, npwp: string, address: string) => {
	const response = await fetch('http://localhost:8080/create_employee', {
		method: 'POST',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ name, npwp, address })
	});

	return response;
};

// src/lib/api_employee.ts

export const getIdEmployee = async (id: string) => {
	const res = await fetch(`http://localhost:8080/update_employee/${id}`);
	if (!res.ok) {
		throw new Error('Gagal mengambil data karyawan');
	}
	return res.json();
};

export const updateEmployee = async (id: string, name: string, npwp: string, address: string) => {
	const res = await fetch(`http://localhost:8080/update_employee/${id}`, {
		method: 'PATCH', // atau PUT, sesuai kebutuhan API
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({ name, npwp, address })
	});

	if (!res.ok) {
		throw new Error('Gagal mengupdate data');
	}
	return res.json();
};
