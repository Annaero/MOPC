import { writable, type Writable } from 'svelte/store';

const today_date: Date = new Date();
export const today: Writable<Date> = writable(today_date);
export const selected_year: Writable<number> = writable(today_date.getFullYear());
export const selected_month: Writable<number> = writable(today_date.getMonth());