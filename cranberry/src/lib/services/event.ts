import dayjs from "dayjs";

import { type Event, EventSchema } from "$lib/models/event";


function parse_event(event_json) {
    const parsed_event = EventSchema.parse(event_json)
    return parsed_event;
};


export async function getEvents(startDay: Date, endDay: Date): Promise<Event[]> {
    /**
     * Get all events happened between two dates
     * 
     * @param startDay {Date} 
     * @param endDay {Date}
     */

    const response = await fetch(
        `/events?${new URLSearchParams({
            start_date: dayjs(startDay).format("YYYY-MM-DD"),
            end_date: dayjs(endDay).format("YYYY-MM-DD"),
        })}`
    );
    const events_json = await response.json();
    const events = events_json.map(parse_event);
    return events;

}