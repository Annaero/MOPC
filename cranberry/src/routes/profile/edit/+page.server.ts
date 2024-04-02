import { error, fail, redirect } from '@sveltejs/kit';
import { message, superValidate } from 'sveltekit-superforms/server';
import { ProfileOptionalDefaultsSchema } from '$lib/models';
import prisma from "$lib/db/prisma";
import type { Actions, PageServerLoad } from '../$types';

export const load: PageServerLoad = async ({ params, locals: { supabase } }) => {
    const user = await supabase.auth.getUser();
    if (!user) {
        throw error(403);
    }

    let profile = prisma.profile.findUnique({ where: { userId: user.data.user.id } });

    const form = await superValidate(profile, ProfileOptionalDefaultsSchema);
    return { form, profile };
};

export const actions: Actions = {
    default: async ({ request, locals: { supabase } }) => {
        const user_data = await supabase.auth.getUser();
        if (!user_data) {
            redirect(303, "/auth")
        }
        const user = user_data.data.user;

        const data = await request.formData()
        const form = await superValidate(data, ProfileOptionalDefaultsSchema);
        if (!form.valid) {
            console.log(form)
            return fail(400, { form });
        }

        if (!form.data.userId) {
            form.data.userId = user.id
        }
        else if (form.data.userId != user.id) {
            throw error(403);
        }

        let profile: Profile;

        try {
            if (!form.data.id) {
                profile = await prisma.profile.create({ data: form.data })
            }
            else {
                let { id: id, ...updateProfile
                } = form.data;
                profile = await prisma.profile.update({ where: { id: id }, data: updateProfile })
            }
        } catch (e) {
            console.error(e);
            return message(form, "Something went wrong :(", {
                status: 500
            });
        }

        throw redirect(300, "/profile/id_" + profile.id);
    }
};