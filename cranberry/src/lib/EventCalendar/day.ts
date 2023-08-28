import type { CalendarEvent } from "./calendarEvent";

export type Day = {
    date: Date,
    today: boolean,
    events: CalendarEvent[]
};

export type Week = {
    startDate: Date,
    endDate: Date,
    events: CalendarEvent[]

};