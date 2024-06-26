import { error } from '@sveltejs/kit';
import type { Event } from '$lib/models/event'
import type { PageServerLoad } from './$types';
import { getEvent } from '$lib/db/events';

export const load: PageServerLoad = async ({ params }) => {
	const eventDocument: Event = await getEvent(params.id)

	console.log(eventDocument)
	if (!eventDocument) throw error(404);

	return {
		event: eventDocument
	};
}
