import { error } from '@sveltejs/kit';
import { getEvent } from '../../../services/event';
import type { MOPCEvent } from '../../../models/mopcEvent';

export async function load({ params }) {
	const event: MOPCEvent = await getEvent(params.eventId);

	if (!event) throw error(404);

	return {
		event
	};
}
