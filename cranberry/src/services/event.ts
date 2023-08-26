import type { CalendarEvent } from "../lib/EventCalendar/calendarEvent";
import { dateToISODateStr } from "../lib/dateUtils";

const EVENTS_ENDPOINT = "/events/events"

export async function getEvents(startDay: Date, endDay: Date): Promise<CalendarEvent[]> {
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
        const e: CalendarEvent = {
            id: x.id,
            name: x.name,
            startDate: new Date(Date.parse(x.startDate)),
            endDate: new Date(Date.parse(x.endDate)),
            active: false,
        };
        e.startDate.setHours(0, 0, 0, 0);
        e.endDate.setHours(0, 0, 0, 0);
        return e;
    });

    return events;

}