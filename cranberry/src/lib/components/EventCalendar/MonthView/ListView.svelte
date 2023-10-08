<script lang="ts">
    import { locale, json } from "svelte-i18n";
    import type { Day, Week } from "../day";
    import type { MOPCEvent } from "$lib/models/mopcEvent";
    import { isSameDate } from "$lib/dateUtils";
    import { getContext } from "svelte";

    let selected_event = getContext("selected_event");
    export let events: MOPCEvent[];

    let today: Date = new Date();
    let year = getContext("selected_year");
    let month = getContext("selected_month");
    let weeks: Week[] = [];

    var firstDayOfMonth: Date;
    var lastDayOfMonth: Date;

    let component: HTMLElement;

    const onScroll = (e: UIEvent) => {
        const element: HTMLElement = e.target as HTMLElement;

        const THRESHOLD = element.clientWidth; //rethink, probably min(element.clientWidth, 100px)

        // move up to much
        while (element.scrollTop < THRESHOLD) {
            //add more earlier event
            // remove old events but not too much
            //      e.g. if not events_to_delete.height ...
        }

        while (
            element.scrollHeight - (element.clientWidth + element.scrollTop) <
            THRESHOLD
        ) {
            //add more later events
            // USE SOME EVENTS CACHE CONTAINER
        }
    };
</script>

<div
    class="flex flex-col border-t border-l bg-base-100 mr-2 overflow-scroll after:bg-gradient-to-r after:from-cyan-500 after:to-blue-500"
    style="height: 100px"
    bind:this={component}
>
    <div>
        <div class="sticky top-0 ...">A</div>
        {#each events as event}
            <div><p>{event.name}</p></div>
        {/each}
    </div>
    <div>
        <div class="sticky top-0">B</div>
        <div>
            <div>
                <strong>Bob Alfred</strong>
            </div>
        </div>
    </div>
</div>
