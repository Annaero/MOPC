import { error } from '@sveltejs/kit';
import { getEvent } from '../../../services/event';
import type { Event } from '../../../lib/models/event';
import { ObjectId } from "mongodb"
import db from '$lib/db/events'

export async function load({ params }) {
	const eventDocument = await db.collection("events")
		.findOne({ _id: new ObjectId(params.eventId) },
			{ project: { "id": "$_id", "_id": 0, name: 1, startDate: 1, endDate: 1 } })

	console.log(eventDocument)
	if (!eventDocument) throw error(404);

	return {
		eventDocument
	};
}
