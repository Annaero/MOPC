import { error, fail, redirect } from '@sveltejs/kit';
import { EventType } from "@prisma/client"

import { message, superValidate } from 'sveltekit-superforms/server';

import { getEvent } from '$lib/db/events';
import prisma from "$lib/db/prisma";
import { type Event, EventOptionalDefaultsSchema } from '$lib/models';


/** @type {import('./$types').PageLoad} */
export async function load({ params, locals: { supabase } }) {
    const user = await supabase.auth.getUser();
    if (!user) {
        throw error(403);
    }

    const id = params.id;
    let event: Event = null;
    if (id) {
        event = await getEvent(params.id)
        if (event.ownerId != user.data.user.id) {
            throw error(403);
        }
    }
    const form = await superValidate(event, EventOptionalDefaultsSchema);
    const eventTypes = Object.values(EventType)
    return { form, eventTypes };
};

/** @type {import('./$types').Actions} */
export const actions = {
    default: async ({ request, locals: { supabase } }) => {
        const user = await supabase.auth.getUser();
        if (!user) {
            throw error(403);
        }

        const data = await request.formData()
        const form = await superValidate(data, EventOptionalDefaultsSchema);
        if (!form.data.ownerId) {
            form.data.ownerId = user.data.user.id
        }

        if (!form.valid) {
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
        } catch (e) {
            console.error(e);
            return message(form, "Something went wrong :(", {
                status: 500
            });
        }

        throw redirect(300, "/events/id_" + event.id);
    }
};