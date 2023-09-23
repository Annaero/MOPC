import { z } from "zod"

export const eventSchema = z.object({
    id: z.number(),
    name: z.string(),
    startDate: z.date(),
    endDate: z.date(),
});
