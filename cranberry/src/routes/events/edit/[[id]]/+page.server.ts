import { fail, error, redirect } from '@sveltejs/kit';
import { message, superValidate } from 'sveltekit-superforms/server';
import { type Event, EventOptionalDefaultsSchema } from '$lib/models/event';
import prisma from "$lib/db/prisma";
import { EventType, Prisma } from "@prisma/client"

import type { Actions, PageServerLoad } from '../$types';

export const load: PageServerLoad = async ({ params }) => {
    let event: Event = null;
    if (params['id']) {
        event = await prisma.event.findUnique({
            where: { id: params['id'] }
        });

        if (!event) throw error(404);
    }

    const form = await superValidate(event, EventOptionalDefaultsSchema);
    const eventTypes = Object.values(EventType)
    console.log("FORM:", form)

    // Always return { form } in load and form actions.
    return { form, eventTypes };
};

export const actions: Actions = {
    default: async ({ request }) => {
        const data = await request.formData()
        console.log('POST:', data);
        const form = await superValidate(data, EventOptionalDefaultsSchema);
        console.log('VALIDATE RESULT:', form);

        // Convenient validation check:
        if (!form.valid) {
            // Again, return { form } and things will just work.
            console.warn("return invalid")
            return fail(400, { form });
        }

        let event: Event;

        try {
            if (!form.data.id) {
                event = await prisma.event.create({ data: form.data })
            }
            else {
                let { id: id, ...updateEvent
                } = form.data;
                event = await prisma.event.update({ where: { id: id }, data: updateEvent })
            }
            console.log("UPSERTED EVENT:", event)
        } catch (e) {
            console.error("THROW ERROR: ", e);
            return message(form, "Something went wrong :(", {
                status: 500
            });
        }

        throw redirect(300, "/events/" + event.id);
    }
};