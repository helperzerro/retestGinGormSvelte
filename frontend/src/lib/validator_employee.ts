import { z } from 'zod';

// Schema validasi Zod
const employeeSchema = z.object({
	name: z.string().min(1, 'Nama tidak boleh kosong'),
	npwp: z
		.string()
		.min(1, 'NPWP tidak boleh kosong')
		.regex(/^\d{15}$/, 'NPWP harus terdiri dari 15 digit angka'),
	address: z.string().min(1, 'Alamat tidak boleh kosong')
});

export const validatorEmployee = async (
	name: string,
	npwp: string,
	address: string,
	setErrors: (errors: Record<string, string>) => void
): Promise<boolean> => {
	// Menambahkan Promise<boolean> agar mengembalikan nilai boolean
	const result = employeeSchema.safeParse({ name, npwp, address });

	if (!result.success) {
		const errors: Record<string, string> = {};
		result.error.errors.forEach((err) => {
			errors[err.path[0]] = err.message;
		});
		setErrors(errors);
		return false; // Mengembalikan false jika validasi gagal
	}
	return true; // Mengembalikan true jika validasi berhasil
};
