<script lang="ts">
    import MonthViewHeader from "./MonthViewHeader.svelte";
    import type { Event } from "$lib/models/event";
    import EventCard from "$lib/components/eventcard/EventCard.svelte";

    export let events: Array<Event>;
</script>

<div
    class="grid xl:grid-cols-3 lg:grid-cols-2 md:grid-cols-1 w-full py-2 space-y-4 justify-around justify-items-center"
>
    {#each events as event}
        <EventCard {event} minified={true} />
    {/each}
    {#if events.length == 0}
        <p>No event this month :(</p>
    {/if}
</div>
