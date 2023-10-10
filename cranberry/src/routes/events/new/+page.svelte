<script lang="ts">
    import EditEvent from "$lib/components/eventcard/EditEvent.svelte";
    import { postEvent } from "../../../services/event";

    import type { MOPCEvent } from "$lib/models/mopcEvent";

    let event: MOPCEvent = {
        id: null,
        name: "",
        startDate: new Date(),
        type: "offline",
    };

    function submitEvent() {
        postEvent(event);
    }
</script>

<EditEvent {event} submitAction={submitEvent} />
