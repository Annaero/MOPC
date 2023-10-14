import type { Event } from "../../models/event";

export type Day = {
    date: Date,
    today: boolean,
    events: Event[]
};

export type Week = {
    startDate: Date,
    endDate: Date,
    events: Event[]

};