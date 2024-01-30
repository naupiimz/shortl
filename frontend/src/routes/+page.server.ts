import type { Actions } from '@sveltejs/kit';

export const actions: Actions = {
	create: async ({ request }) => {
		const data = await request.formData();
		const url = data.get('long_url');
		const id = data.get('user_id');

		const response = await fetch('http://localhost:9000/create-short-url', {
			method: 'POST',
			headers: {
				'content-type': 'application/json'
			},
			body: JSON.stringify({ long_url: url, user_id: id })
		});

		const result = await response.json();

		if (response.ok) {
			return { result, success: true };
		} else {
			return { success: false };
		}
		return { success: false };
	}
};
