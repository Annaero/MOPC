import prisma from "$lib/db/prisma";
import { error } from "@sveltejs/kit";

export const load = async ({ locals: { getSession } }) => {
    const session = await getSession()
    const userCard: string = null;
    if (session) {
        const userId = session.user.id
        const userCard = await prisma.userCard.findFirst({ where: { userId: userId } })

        // if (userCard == null) {
        //     console.error(`User ${userId} does not have a userCard assosiated`)
        //     throw error(500)
        // }
    }

    return {
        session: session,
        user: userCard
    }
}