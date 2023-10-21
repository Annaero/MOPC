import { error } from '@sveltejs/kit';
import type { Event } from '$lib/models/event'
import prisma from '$lib/db/prisma'

export async function load({ params }) {
	const eventDocument: Event = await prisma.event.findUnique({
		where: { id: params.eventId }
	});

	console.log(eventDocument)
	if (!eventDocument) throw error(404);

	return {
		event: eventDocument
	};
}
