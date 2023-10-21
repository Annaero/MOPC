<script>
    import EventCard from "$lib/components/eventcard/EventCard.svelte";
    export let data;
</script>

<EventCard event={data.event} />
