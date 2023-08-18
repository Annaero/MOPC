<script lang="ts">
    import MonthView from "./MonthView/MonthView.svelte";
    import { setContext, createEventDispatcher, onDestroy } from "svelte";
    import { writable } from "svelte/store";
    import type { CalendarEvent } from "./calendarEvent";
    const dispatch = createEventDispatcher();

    export let showYear: number;
    export let showMonth: number;

    let selected_event = writable(-1);
    let selected_year = writable(showYear);
    let selected_month = writable(showMonth);

    const unsubscribeSelectedYear = selected_month.subscribe((value) => {
        dispatch("selectedDateChanged", {
            year: selected_year,
            selected_month: value,
        });
    });

    const unsubscribeSelectedMonth = selected_month.subscribe((value) => {
        dispatch("selectedDateChanged", {
            year: selected_year,
            selected_month: value,
        });
    });

    onDestroy(() => {
        unsubscribeSelectedYear();
        unsubscribeSelectedMonth();
    });

    setContext("selected_event", selected_event);
    setContext("selected_year", selected_year);
    setContext("selected_month", selected_month);

    export let events: CalendarEvent[] = [];
    events.sort((e1, e2) => e1.startDate.getTime() - e2.startDate.getTime());
    setContext("events", events);
</script>

<MonthView />
