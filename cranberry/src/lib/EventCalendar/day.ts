import type { MOPCEvent } from "../../models/mopcEvent";

export type Day = {
    date: Date,
    today: boolean,
    events: MOPCEvent[]
};

export type Week = {
    startDate: Date,
    endDate: Date,
    events: MOPCEvent[]

};