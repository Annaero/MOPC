import db from "$lib/db/events"

export async function GET(request: Request) {
    const url = new URL(request.url);
    let startDay = new Date(url.searchParams.get("start_date"));
    let endDay = new Date(url.searchParams.get("end_date"));

    const events = await db.collection("events").find({
        startDate: { $lte: endDay },
        endDate: { $gte: startDay }
    }).project({ "id": "$_id", "_id": 0, name: 1, startDate: 1, endDate: 1 }).toArray();
    console.log(events)

    return new Response(JSON.stringify(events));
}