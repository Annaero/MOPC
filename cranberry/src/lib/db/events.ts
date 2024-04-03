import { error } from '@sveltejs/kit';
import type { Event } from '$lib/models/event'
import prisma from '$lib/db/prisma'
import { Prisma } from '@prisma/client';

export const getEvent = async (id: number) => {
    let eventDocument: Event;
    try {
        eventDocument = await prisma.event.findUnique({
            where: { id: parseInt(id) }
        });
    }
    catch (e) {
        if (e instanceof Prisma.PrismaClientKnownRequestError) {
            if (e.code === 'P2023') {
                console.error((e as Prisma.PrismaClientKnownRequestError).meta?.message)
                throw error(404)
            }
        }
        console.error(e)
        throw e
    }

    return eventDocument
}