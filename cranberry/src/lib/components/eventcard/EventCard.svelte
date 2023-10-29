<script lang="ts">
    import type { Event } from "$lib/models/event";
    import { _ } from "svelte-i18n";
    import DateSheet from "../datesheet/DateSheet.svelte";
    import Icon from "@iconify/svelte";
    import { goto } from "$app/navigation";

    export let minified: boolean = false;
    export let event: Event;
    export let redirect: string = `/events/id_${event.id}`;
</script>

<div
    class="card max-w-xl bg-base-100 shadow-xl border border-solid border-neutral hover:border-base-300 hoer:shadow-inner hover:cursor-default"
    class:w-64={minified}
    tabindex="0"
    role="button"
    on:keyup={() => goto(redirect)}
    on:click={() => goto(redirect)}
>
    {#if !minified}
        <div class="navbar pt-1 px-3">
            <a class="navbar-start" href="/events/edit/id_{event.id}">
                <Icon icon="carbon:edit" height="2em" width="2em" />
            </a>
            <div class="badge badge-success navbar-center">active</div>
            <a class="navbar-end">
                <Icon icon="carbon:star" height="2em" width="2em" />
            </a>
        </div>
    {/if}
    <figure class="px-3 pt-1">
        <img
            src="https://sun9-19.userapi.com/impf/tmr-b5CVibMPy3RiAm_biqiYSJiZEK2TthGeCg/rV0SjdKtsis.jpg?size=1920x768&quality=95&crop=0,639,2396,957&sign=3d064ce9265c2649cd8934b7e40edf99&type=cover_group"
            alt="Shoes"
            class="rounded-xl"
        />
    </figure>
    <div
        class="items-center text-center"
        class:flex={minified}
        class:flex-col={minified}
        class:p-2={minified}
        class:card-body={!minified}
    >
        <h1 class="card-title text-4xl" class:text-xl={minified}>
            {event.name}
        </h1>
        <div class="flex flex-row items-center text-center">
            {#if event.type == "ONLINE"}
                <Icon icon="carbon:screen" />
            {:else if event.type == "OFFLINE"}
                <Icon icon="carbon:location" />
            {/if}
            <p class="text-xs">{event.type}</p>
        </div>
        <div
            class="pt-2 italic pb-2 hyphens-auto"
            class:line-clamp-6={minified}
        >
            {#each event.description.split(/\r?\n/) as paragraph}
                <p
                    class="prose text-left pt-1 indent-5 hyphens-auto"
                    class:text-sm={minified}
                >
                    {paragraph}
                </p>
            {/each}
        </div>
        {#if !minified}
            <div class="divider"><p class="text-neutral">When</p></div>
        {/if}
        <div class="w-5/6">
            <DateSheet
                startDate={event.startDate}
                endDate={event.endDate}
                {minified}
            />
        </div>
        {#if !minified}
            <div class="card-actions pt-10">
                <button class="btn btn-primary">Apply!</button>
            </div>
        {/if}
    </div>
</div>
