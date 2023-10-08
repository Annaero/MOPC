<script lang="ts">
    import EditEvent from "$lib/components/eventcard/EditEvent.svelte";

    import type { MOPCEvent } from "$lib/models/mopcEvent";

    let event: MOPCEvent = { id: null, name: "", startDate: new Date() };

    function postEvent() {
        console.log(event);
    }
</script>

<EditEvent {event} submitAction={postEvent} />
