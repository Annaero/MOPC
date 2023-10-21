import type { Event } from '$lib/models/event';
import prisma from '$lib/db/prisma';

export async function GET(request: Request) {
    const url = new URL(request.url);
    let startDay = new Date(url.searchParams.get("start_date"));
    let endDay = new Date(url.searchParams.get("end_date"));

    const events: Array<Event> = await prisma.event.findMany({
        where: {
            OR: [
                {
                    startDate: { lte: endDay },
                    endDate: { gte: startDay }
                },
                {
                    startDate: { gte: startDay, lte: endDay },
                }

            ]
        }
    })
    console.log(events)

    return new Response(JSON.stringify(events));
}