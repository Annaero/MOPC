import { fail, redirect } from '@sveltejs/kit';
import { message, superValidate } from 'sveltekit-superforms/server';
import type { Actions, PageServerLoad } from '../$types';
import { loginShema } from '../../lib/models/login';

export const load: PageServerLoad = async ({ params }) => {
    const form = await superValidate(loginShema);
    return { form };
};

export const actions: Actions = {
    default: async ({ request }) => {
        const data = await request.formData()
        const form = await superValidate(data, loginShema);

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
            return message(form, "Something went wrong :(", {
                status: 500
            });
        }

        throw redirect(300, "/events/id_" + event.id);
    }
};