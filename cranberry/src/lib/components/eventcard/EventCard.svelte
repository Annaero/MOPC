<script lang="ts">
    import type { Event } from "$lib/models/event";
    import { _ } from "svelte-i18n";
    import DateSheet from "../datesheet/DateSheet.svelte";
    import Icon from "@iconify/svelte";

    export let event: Event;
</script>

<div
    class="card max-w-md bg-base-100 shadow-xl border border-solid border-base-200"
>
    <div class="navbar pt-1 px-3">
        <a class="navbar-start" href="/events/edit/id_{event.id}">
            <Icon icon="carbon:edit" height="2em" width="2em" />
        </a>
        <div class="badge badge-success navbar-center">active</div>
        <a class="navbar-end">
            <Icon icon="carbon:star" height="2em" width="2em" />
        </a>
    </div>
    <figure class="px-3 pt-1">
        <img
            src="https://sun9-19.userapi.com/impf/tmr-b5CVibMPy3RiAm_biqiYSJiZEK2TthGeCg/rV0SjdKtsis.jpg?size=1920x768&quality=95&crop=0,639,2396,957&sign=3d064ce9265c2649cd8934b7e40edf99&type=cover_group"
            alt="Shoes"
            class="rounded-xl"
        />
    </figure>
    <div class="card-body items-center text-center">
        <h1 class="card-title">
            {event.name}
        </h1>
        <div class="flex flex-row items-center text-center">
            {#if event.type == "ONLINE"}
                <Icon icon="carbon:screen" />
            {:else if event.type == "OFFLINE"}
                <Icon icon="carbon:location" />
            {/if}
            <p>{event.type}</p>
        </div>
        <p class="pt-2">{event.description}</p>
        <div class="divider pt-5"><p class="text-neutral">When</p></div>
        <div>
            <DateSheet startDate={event.startDate} endDate={event.endDate} />
        </div>
        <div class="card-actions pt-10">
            <button class="btn btn-primary">Apply!</button>
        </div>
    </div>
</div>
