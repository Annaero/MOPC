import { z } from "zod"
import mongoose from 'mongoose';
const { Schema } = mongoose;


export const EVENT_TYPES = ['online', 'offline', 'duel'] as const

export type Event = {
    id?: string,
    name: string
    description: string,
    startDate: Date,
    endDate?: Date,
    type: string
    active?: boolean,
};

export const eventSchemaZ = z.object({
    id: z.number().nullable(),
    name: z.string().min(5, "Name should be at least 5 character long"),
    description: z.string(),
    type: z.enum(EVENT_TYPES),
    startDate: z.date(),
    endDate: z.date().nullable(),
}).refine(schema => schema.endDate == null || schema.endDate >= schema.startDate,
    "End date mast be same or later than start date");

export const eventSchemaM = new Schema(
    {
        name: String,
        description: String,
        type: EVENT_TYPES,
        startDate: Date,
        endDate: Date
    },
    { timestamps: true });

export const EventsDB = mongoose.model('Event', eventSchemaM);