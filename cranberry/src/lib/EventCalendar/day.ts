import type { CalendarEvent } from "./calendarEvent";

export type Day = {
    date: Date,
    today: boolean,
    events: CalendarEvent[]
};