import type { LayoutServerLoad } from './$types'
import prisma from '$lib/db/prisma';
import { redirect } from '@sveltejs/kit';

/** @type {import('./$types').LayoutServerLoad} */
export async function load({ locals: { safeGetSession } }) {
    let session = await safeGetSession();

    const profile = session ? await prisma.profile.findUnique({
        where: {
            userId: session.user.id,
        }
    }) : undefined;

    return {
        session: session,
        profile: profile
    };
};