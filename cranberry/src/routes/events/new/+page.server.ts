import { fail } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms/server';
import type { Actions, PageServerLoad } from './$types';
import { EventOptionalDefaultsSchema } from '$lib/models/event';
import prisma from "$lib/db/prisma";
import { EventType } from "@prisma/client"

const schema = EventOptionalDefaultsSchema

export const load: PageServerLoad = async () => {
    // Server API:
    const form = await superValidate(schema);
    const eventTypes = Object.values(EventType)

    // Always return { form } in load and form actions.
    return { form, eventTypes };
};

export const actions: Actions = {
    default: async ({ request }) => {
        const form = await superValidate(request, schema);
        console.log('POST:', form);

        // Convenient validation check:
        if (!form.valid) {
            // Again, return { form } and things will just work.
            console.log("return iinvalid")
            return fail(400, { form });
        }

        const createEvent = await prisma.event.create({ data: form.data })
        console.log("w8t", createEvent)
        // .then(() => {
        //     console.log('test')
        // })
        // .catch((err) => {
        //     console.log(`Error saving Data. 'Error:`, err)
        // });

        // Unless you throw, always return { form } in load and form actions.
        console.log("returb")
        return { form };
    }
};