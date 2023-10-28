<script lang="ts">
    import EventCalendar from "$lib/components/EventCalendar/EventCalendar.svelte";
    import { getEvents } from "$lib/services/event";
    import type { Event } from "$lib/models/event";

    let today: Date = new Date();
    let yearToShow: number = today.getFullYear();
    let monthToShow: number = today.getMonth();

    let events: Event[] = [];

    async function updateEvents(year: number, month: number) {
        const firstayOfMonth = new Date(year, month, 1);
        const lastDayOfMonth = new Date(year, month + 1, 0);
        events = await getEvents(firstayOfMonth, lastDayOfMonth);
    }
</script>

<div class="w-full">
    {#await updateEvents(yearToShow, monthToShow)}
        <span class="loading loading-spinner loading-lg" />
    {:then _value}
        <EventCalendar
            showYear={yearToShow}
            showMonth={monthToShow}
            {events}
            on:selectedDateChanged={async (e) => {
                await updateEvents(e.detail.year, e.detail.month);
            }}
        />
    {/await}
</div>
