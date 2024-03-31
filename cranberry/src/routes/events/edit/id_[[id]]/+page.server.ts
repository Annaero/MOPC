import { error, fail, redirect } from '@sveltejs/kit';
import { message, superValidate } from 'sveltekit-superforms/server';
import { type Event, EventOptionalDefaultsSchema } from '$lib/models';
import prisma from "$lib/db/prisma";
import { EventType } from "@prisma/client"
import type { Actions, PageServerLoad } from '../$types';
import { getEvent } from '$lib/db/events';

export const load: PageServerLoad = async ({ params, locals: { supabase } }) => {
    // const session = await getSession();
    const user = supabase.auth.getUser();
    if (!user) {
        throw error(403);
    }

    const id = params.id;
    let event: Event = null;
    if (id) {
        event = await getEvent(params.id)
        if (event.ownerId != user.id) {
            throw error(403);
        }
    }
    const form = await superValidate(event, EventOptionalDefaultsSchema);
    const eventTypes = Object.values(EventType)
    return { form, eventTypes };
};

export const actions: Actions = {
    default: async ({ request, locals: { supabase } }) => {
        const user = supabase.auth.getUser();
        if (!user) {
            throw error(403);
        }

        const data = await request.formData()
        const form = await superValidate(data, EventOptionalDefaultsSchema);
        if (!form.data.ownerId) {
            form.data.ownerId = user.id
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