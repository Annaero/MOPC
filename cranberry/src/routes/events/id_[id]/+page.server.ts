import { error } from '@sveltejs/kit';
import type { Event } from '$lib/models'
import type { PageServerLoad } from './$types';
import { getEvent } from '$lib/db/events';

export const load: PageServerLoad = async ({ params }) => {
	const eventDocument: Event = await getEvent(params.id)

	if (!eventDocument) throw error(404);

	return {
		event: eventDocument
	};
}
