<script lang="ts">
    import { _ } from "svelte-i18n";
    import { getContext } from "svelte";
    import Icon from "@iconify/svelte";

    let year = getContext("selected_year");
    let month = getContext("selected_month");

    export let viewOptions: Array<String>;
    console.log(viewOptions);
    export let selectedView = "calendar"; //viewOptions[0];

    function moveMonth(n: number) {
        let nextMonthDate: Date = new Date($year, $month + n, 1);
        $year = nextMonthDate.getFullYear();
        $month = nextMonthDate.getMonth();
    }
</script>

<div
    class="navbar bg-base-100 rounded-box w-full drop-shadow-md border border-solid border-base-200"
>
    <div class="navbar-start">
        <span class="text-lg font-bold">
            {$_("date.month_names")[$month]}
        </span>
        <span class="ml-1 text-lg font-normal">{$year}</span>
    </div>
    <div class="navbar-center">
        <div class="join" role="radiogroup">
            {#each viewOptions as view}
                <input
                    type="radio"
                    name="options"
                    aria-label={view}
                    value={view}
                    bind:group={selectedView}
                    class="join-item btn btn-outline btn-sm border-base-200 min-w-15"
                />
            {/each}
        </div>
    </div>
    <div class="navbar-end">
        <div class="join">
            <button
                type="button"
                on:click={(e) => {
                    moveMonth(-1);
                }}
                class="join-item btn btn-outline border-base-200"
            >
                <Icon icon="carbon:chevron-left" height="2em" width="2em" />
            </button>
            <!-- <div class="border-r inline-flex h-6" /> -->
            <button
                type="button"
                on:click={(e) => {
                    moveMonth(1);
                }}
                class="join-item btn btn-outline border-base-200"
            >
                <Icon icon="carbon:chevron-right" height="2em" width="2em" />
            </button>
        </div>
    </div>
</div>
