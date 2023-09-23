<script>
    import EventCard from "$lib/components/eventcard/EventCard.svelte";

    export let data;
</script>

<EventCard morcEvent={data.event} />
