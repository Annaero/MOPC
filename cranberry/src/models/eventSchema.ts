import { z } from "zod"

const eventSchema = z.object({
    id: z.number(),
    name: z.string(),
    startDate: z.date(),
    endDate: z.date(),
});
