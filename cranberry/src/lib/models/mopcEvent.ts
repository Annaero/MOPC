import { z } from "zod"

export type MOPCEvent = {
    id?: string,
    name: string
    description: string,
    startDate: Date,
    endDate?: Date,
    type: string
    active?: boolean,
};

export const eventSchema = z.object({
    id: z.number().nullable(),
    name: z.string().min(5, "Name should be at least 5 character long"),
    description: z.string(),
    startDate: z.date(),
    endDate: z.date().nullable(),
}).refine(schema => schema.endDate == null || schema.endDate >= schema.startDate,
    "End date mast be same or later than start date");