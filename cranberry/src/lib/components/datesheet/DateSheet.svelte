<script lang="ts">
    import Icon from "@iconify/svelte";
    import { _ } from "svelte-i18n";

    export let startDate: Date;
    export let endDate: Date = null;
</script>

<div class="flex flex-row justify-center w-96">
    <div
        class="flex flex-col border border-solid border-base-200 rounded-box w-1/4"
    >
        <div class="text-l bg-base-200 rounded-t-lg w-full">
            {$_("date.month_names")[startDate.getMonth()]}
        </div>
        <div class="text-2xl">{startDate.getDate()}</div>
        <div class="text-m">
            {startDate.getFullYear()}
        </div>
    </div>
    {#if endDate}
        <div class="flex flex-row h-full place-self-center w-1/8 px-2">
            <p class="text-xl text-center">—</p>
        </div>
        <div
            class="flex flex-col place-items-center border border-solid border-base-200 rounded-box w-1/4"
        >
            <div class="text-l bg-base-200 rounded-t-lg w-full">
                {$_("date.month_names")[endDate.getMonth()]}
            </div>
            <div class="text-2xl">{endDate.getDate()}</div>
            <div class="text-m">
                {endDate.getFullYear()}
            </div>
        </div>
    {/if}
</div>
