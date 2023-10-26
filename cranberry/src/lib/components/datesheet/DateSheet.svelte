<script lang="ts">
    import Icon from "@iconify/svelte";
    import { _ } from "svelte-i18n";

    export let minified: boolean = false;

    export let startDate: Date;
    export let endDate: Date = null;
</script>

<div class="flex flex-row justify-center w-full">
    <div
        class="flex flex-col border border-solid border-base-200 rounded-box w-1/3"
    >
        <div
            class="text-l bg-base-200 rounded-t-lg w-full"
            class:text-m={minified}
        >
            {startDate.getFullYear()}
        </div>
        <div class="text-2xl" class:text-xl={minified}>
            {startDate.getDate()}
        </div>
        <div class="text-m" class:text-xs={minified}>
            {$_("date.month_names")[startDate.getMonth()]}
        </div>
    </div>
    {#if endDate}
        <div class="flex flex-row h-full place-self-center w-1/8 px-2">
            <p class="text-xl text-center">—</p>
        </div>
        <div
            class="flex flex-col place-items-center border border-solid border-base-200 rounded-box w-1/3"
        >
            <div class="text-l bg-base-200 rounded-t-lg w-full">
                {endDate.getFullYear()}
            </div>
            <div class="text-2xl" class:text-xl={minified}>
                {endDate.getDate()}
            </div>
            <div class="text-m" class:text-xs={minified}>
                {$_("date.month_names")[endDate.getMonth()]}
            </div>
        </div>
    {/if}
</div>
