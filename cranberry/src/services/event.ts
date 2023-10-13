import dayjs from "dayjs";

import type { MOPCEvent } from "../lib/models/mopcEvent";
import { dateToISODateStr } from "../lib/dateUtils";

const EVENTS_ENDPOINT = "/api/events"
const format = "YYYY-MM-DD";

function parse_event(event_json) {
    const e: MOPCEvent = {
        id: event_json.id,
        name: event_json.name,
        description: event_json.description,
        startDate: new Date(event_json.startDate),
        endDate: event_json.endDate != "" ? new Date(event_json.endDate) : null,
        active: false,
    };
    e.startDate.setHours(0, 0, 0, 0);
    if (e.endDate !== null) {
        e.endDate.setHours(0, 0, 0, 0);
    }

    return e;
};

function dateReplacer(key, value) {
    if (this[key] instanceof Date) {
        return stringifyDate(this[key]);
    }
    return value ?? undefined;
}

export async function getEvents(startDay: Date, endDay: Date): Promise<MOPCEvent[]> {
    /**
     * Get all events happened between two dates
     * 
     * @param startDay {Date} 
     * @param endDay {Date}
     */

    const response = await fetch(
        `events/?${new URLSearchParams({
            start_date: dayjs(startDay).format("YYYY-MM-DD"),
            end_date: dayjs(endDay).format("YYYY-MM-DD"),
        })}`
    );
    const events_json = await response.json();
    const events = events_json.map(parse_event);
    return events;

}

export async function getEvent(eventID): Promise<MOPCEvent> {
    /**
     * Get event by its id
     * 
     * @param eventID {number}
     */

    const response = await fetch(EVENTS_ENDPOINT + "/" + eventID);
    const event_json = await response.json();
    return parse_event(event_json);

}


function stringifyDate(date: Date) {
    return dayjs(date).format(format)
}

export async function postEvent(event: MOPCEvent): Promise<string> {
    /**
     * Get all events happened between two dates
     * 
     * @param startDay {Date} 
     * @param endDay {Date}
     */
    const eventJson = JSON.stringify(event, dateReplacer)

    const response = await fetch(EVENTS_ENDPOINT, {
        method: "POST", headers: {
            "Content-Type": "application/json",
        },
        body: eventJson
    })

    const responseJson = await response.json()
    if (!response.ok) {
        throw new Error(responseJson)
    }

    return responseJson["id"]
}