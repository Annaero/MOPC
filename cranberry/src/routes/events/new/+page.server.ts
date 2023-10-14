import { fail } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms/server';
import { eventSchemaZ, EventsDB } from '$lib/models/event';

import db from '$lib/db/events'

import type { Actions, PageServerLoad } from './$types';

const schema = eventSchemaZ

export const load: PageServerLoad = async () => {
    // Server API:
    const form = await superValidate(schema);

    // Always return { form } in load and form actions.
    return { form };
};

export const actions: Actions = {
    default: async ({ request }) => {
        const form = await superValidate(request, schema);
        console.log('POST', form);

        // Convenient validation check:
        if (!form.valid) {
            // Again, return { form } and things will just work.
            return fail(400, { form });
        }

        const db_event = new EventsDB(form.data)
        await db_event.save()
            .then(() => {
                console.log('test')
            })
            .catch((err) => {
                console.log(`Error saving Data. 'Error:`, err)
            });

        // Unless you throw, always return { form } in load and form actions.
        return { form };
    }
};