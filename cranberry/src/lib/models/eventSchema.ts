import { z } from "zod"

export const eventSchema = z.object({
    id: z.number().nullable(),
    name: z.string().min(5),
    description: z.string(),
    startDate: z.date(),
    endDate: z.date().nullable(),
}).refine(schema => schema.endDate == null || schema.endDate >= schema.startDate);
