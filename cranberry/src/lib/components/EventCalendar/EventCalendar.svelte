<script lang="ts">
    import MonthView from "./MonthView/MonthView.svelte";
    import { setContext, createEventDispatcher, onDestroy } from "svelte";
    import { writable } from "svelte/store";
    import type { MOPCEvent } from "$lib/models/mopcEvent";
    const dispatch = createEventDispatcher();

    export let showYear: number;
    export let showMonth: number;

    let selected_event = writable(-1);
    let selected_year = writable(showYear);
    let selected_month = writable(showMonth);

    const unsubscribeSelectedYear = selected_month.subscribe((value) => {
        dispatch("selectedDateChanged", {
            year: $selected_year,
            month: value,
        });
    });

    const unsubscribeSelectedMonth = selected_year.subscribe((value) => {
        dispatch("selectedDateChanged", {
            year: value,
            month: $selected_month,
        });
    });

    onDestroy(() => {
        unsubscribeSelectedYear();
        unsubscribeSelectedMonth();
    });

    setContext("selected_event", selected_event);
    setContext("selected_year", selected_year);
    setContext("selected_month", selected_month);

    export let events: MOPCEvent[] = [];
    events.sort((e1, e2) => e1.startDate.getTime() - e2.startDate.getTime());

    let eventsCache: Map<number, MOPCEvent[]> = new Map<number, MOPCEvent[]>();

    function eventCacheLookup(
        year: number,
        month: number,
        missCallback: CallableFunction
    ) {
        const key = year * 100 + month;
        if (eventsCache.has(key)) {
            return eventsCache[key];
        }
        missCallback(year, month);
    }
</script>

<MonthView {events} />
