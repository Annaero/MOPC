import { fail, error } from '@sveltejs/kit';
import { message, superValidate } from 'sveltekit-superforms/server';
import { type Event, EventOptionalDefaultsSchema } from '$lib/models/event';
import prisma from "$lib/db/prisma";
import { EventType, Prisma } from "@prisma/client"

import type { Actions, PageServerLoad } from '../$types';

const schema = EventOptionalDefaultsSchema

export const load: PageServerLoad = async ({ params }) => {
    let event: Event = null;
    if (params['id']) {
        event = await prisma.event.findUnique({
            where: { id: params['id'] }
        });

        if (!event) throw error(404);
    }

    const form = await superValidate(event, schema);
    const eventTypes = Object.values(EventType)
    console.log("FORM:", form)

    // Always return { form } in load and form actions.
    return { form, eventTypes };
};

export const actions: Actions = {
    default: async ({ request }) => {
        const data = await request.formData()
        console.log('POST:', data);
        const form = await superValidate(data, schema);
        console.log('VALIDATE RESULT:', form);

        // Convenient validation check:
        if (!form.valid) {
            // Again, return { form } and things will just work.
            console.warn("return invalid")
            return fail(400, { form });
        }

        let upsertEvent: Event;
        try {
            let { id: _, ...updateEvent
            } = form.data;
            upsertEvent = await prisma.event.upsert({
                where: { id: form.data.id },
                update: updateEvent,
                create: form.data
            })
            console.log("UPSERTED EVENT:", upsertEvent)
            return { form };
        } catch (e) {
            console.error("THROW ERROR: ", e);
            return message(form, JSON.stringify(e), {
                status: 500
            });
        }
    }
};