<script lang="ts">
    import { setContext, createEventDispatcher, onDestroy } from "svelte";
    import { writable } from "svelte/store";
    import type { Event } from "$lib/models/event";
    import MonthViewHeader from "./MonthView/MonthViewHeader.svelte";
    import MonthViewCalendar from "./MonthView/MonthViewCalendar.svelte";
    import MonthViewCards from "./MonthView/MonthViewCards.svelte";
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

    export let events: Event[] = [];
    events.sort((e1, e2) => e1.startDate.getTime() - e2.startDate.getTime());

    const viewOptions = ["calendar", "overview"];
    let selectedView: string = "calendar";
</script>

<MonthViewHeader {viewOptions} bind:selectedView />
{#if selectedView === viewOptions[0]}
    <MonthViewCalendar {events} />
{:else}
    <MonthViewCards {events} />
{/if}
