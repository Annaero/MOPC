import type { CalendarEvent } from "./event";

export type Day = {
    date: Date,
    today: boolean,
    events: CalendarEvent[]
};