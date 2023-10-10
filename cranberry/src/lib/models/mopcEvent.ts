export type MOPCEvent = {
    id?: string,
    name: string
    description: string,
    startDate: Date,
    endDate?: Date,
    type: string
    active?: boolean,
};