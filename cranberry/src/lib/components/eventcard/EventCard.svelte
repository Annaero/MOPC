<script lang="ts">
    import { json } from "svelte-i18n";
    import type { MOPCEvent } from "$lib/models/mopcEvent";

    export let morcEvent: MOPCEvent;
    export let showInline: boolean = false;
</script>

<div class="flex flex-col w-full bg-white rounded shadow-lg">
    {#if !showInline}
        <div
            class="w-full h-64 bg-top bg-cover rounded-t"
            style="background-position: center; background-image: url(https://sun9-19.userapi.com/impf/tmr-b5CVibMPy3RiAm_biqiYSJiZEK2TthGeCg/rV0SjdKtsis.jpg?size=1920x768&quality=95&crop=0,639,2396,957&sign=3d064ce9265c2649cd8934b7e40edf99&type=cover_group)"
        />
    {/if}
    <div class="flex flex-col w-full md:flex-row">
        <div
            class="flex flex-row justify-around p-4 font-bold leading-none text-gray-800 uppercase bg-red-100 rounded md:flex-col md:items-center md:justify-center md:w-1/4"
        >
            <div class="md:text-6xl">{morcEvent.startDate.getDate()}</div>
            <div class="md:text-xl">
                {$json("date.month_names")[morcEvent.startDate.getMonth()]}
            </div>
            <div class="md:text-2xl">{morcEvent.startDate.getFullYear()}</div>
        </div>
        <div class="p-4 font-normal text-gray-800 md:w-3/4">
            <h1
                class="mb-4 text-4xl font-bold leading-none tracking-tight text-gray-800"
            >
                {morcEvent.name}
            </h1>
            <p class="leading-normal">
                {morcEvent.name}
            </p>
            <div class="flex flex-row items-center mt-4 text-gray-700">
                <div class="w-1/2 flex justify-end">
                    <img
                        src="https://collegefootballplayoff.com/images/section_logo.png"
                        alt=""
                        class="w-8"
                    />
                </div>
            </div>
        </div>
    </div>
</div>
