<script lang="ts">
    import { locale, json } from "svelte-i18n";
    import { getContext } from "svelte";

    locale.set("en");
    locale.subscribe(() => console.log("locale change"));

    let year = getContext("selected_year");
    let month = getContext("selected_month");

    function moveMonth(n: number) {
        let nextMonthDate: Date = new Date($year, $month + n, 1);
        $year = nextMonthDate.getFullYear();
        $month = nextMonthDate.getMonth();
    }
</script>

<div
    class="menu bg-base-100 rounded-box w-full drop-shadow-md border border-base-200"
>
    <div class="flex items-center justify-between">
        <div class="flex-1">
            <span class="text-lg font-bold text-bg-neutral">
                {$json("date.month_names")[$month]}
            </span>
            <span class="ml-1 text-lg text-bg-neutral font-normal">
                {$year}</span
            >
        </div>
        <div class="flex-none">
            <div class="join grid grid-cols-2">
                <button
                    type="button"
                    on:click={(e) => {
                        moveMonth(-1);
                    }}
                    class="join-item btn btn-outline border-neutral-content"
                >
                    <svg
                        class="h-6 w-6 text-neutral-content inline-flex leading-none"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M15 19l-7-7 7-7"
                        />
                    </svg>
                </button>
                <!-- <div class="border-r inline-flex h-6" /> -->
                <button
                    type="button"
                    on:click={(e) => {
                        moveMonth(1);
                    }}
                    class="join-item btn btn-outline border-neutral-content"
                >
                    <svg
                        class="h-6 w-6 text-neutral-content inline-flex leading-none"
                        fill="none"
                        viewBox="0 0 24 24"
                        stroke="currentColor"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M9 5l7 7-7 7"
                        />
                    </svg>
                </button>
            </div>
        </div>
    </div>
</div>
