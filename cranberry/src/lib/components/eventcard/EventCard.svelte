<script lang="ts">
    import type { Event } from "$lib/models/event";
    import { _ } from "svelte-i18n";

    export let event: Event;
    export let showInline: boolean = false;
</script>

<div class="flex flex-col w-full rounded shadow-lg">
    {#if !showInline}
        <div
            class="w-full h-64 bg-top bg-cover rounded-t"
            style="background-position: center; background-image: url(https://sun9-19.userapi.com/impf/tmr-b5CVibMPy3RiAm_biqiYSJiZEK2TthGeCg/rV0SjdKtsis.jpg?size=1920x768&quality=95&crop=0,639,2396,957&sign=3d064ce9265c2649cd8934b7e40edf99&type=cover_group)"
        />
    {/if}
    <div class="flex flex-col w-full md:flex-row">
        <div
            class="flex flex-row justify-around p-4 md:flex-col md:items-center md:justify-center md:w-1/4"
        >
            <div class="md:text-6xl">{event.startDate.getDate()}</div>
            <div class="md:text-xl">
                {$_("date.month_names")[event.startDate.getMonth()]}
            </div>
            <div class="md:text-2xl">{event.startDate.getFullYear()}</div>
        </div>
        <div class="p-4 font-normal text-gray-800 md:w-3/4">
            <h1
                class="mb-4 text-4xl font-bold leading-none tracking-tight text-gray-800"
            >
                {event.name}
            </h1>
            <p class="leading-normal">
                {event.description}
            </p>
            <div class="flex flex-row items-center mt-4 text-gray-700">
                <div class="w-1/2 flex justify-end">
                    <a href="/events/edit/{event.id}">
                        <svg
                            class="feather feather-edit"
                            fill="none"
                            height="24"
                            stroke="currentColor"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            viewBox="0 0 24 24"
                            width="24"
                            xmlns="http://www.w3.org/2000/svg"
                            ><path
                                d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"
                            /><path
                                d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"
                            />
                        </svg>
                    </a>
                </div>
            </div>
        </div>
    </div>
</div>
