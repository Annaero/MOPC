import type { MOPCEvent } from "../models/mopcEvent";
import { dateToISODateStr } from "../lib/dateUtils";

const EVENTS_ENDPOINT = "/events/events"

export async function getEvents(startDay: Date, endDay: Date): Promise<MOPCEvent[]> {
    /**
     * Get all events happened between two dates
     * 
     * @param startDay {Date} 
     * @param endDay {Date}
     */

    const response = await fetch(
        EVENTS_ENDPOINT + "?" +
        new URLSearchParams({
            start_date: dateToISODateStr(startDay.getFullYear(), startDay.getMonth(), startDay.getDate()),
            end_date: dateToISODateStr(endDay.getFullYear(), endDay.getMonth(), endDay.getDate()),
        })
    );
    const events_json = await response.json();
    const events = events_json.map((x) => {
        const e: MOPCEvent = {
            id: x._id,
            name: x.name,
            startDate: new Date(Date.parse(x.startDate)),
            endDate: new Date(Date.parse(x.endDate)),
            active: false,
        };
        e.startDate.setHours(0, 0, 0, 0);
        // e.startDate.setMonth(e.startDate.getMonth() - 1);
        e.endDate.setHours(0, 0, 0, 0);
        // e.endDate.setMonth(e.endDate.getMonth() - 1);
        return e;
    });

    return events;

}