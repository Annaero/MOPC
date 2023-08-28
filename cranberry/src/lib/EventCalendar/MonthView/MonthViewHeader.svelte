<script lang="ts">
    import { locale, json } from "svelte-i18n";
    import { getContext } from "svelte";

    locale.set("en");
    locale.subscribe(() => console.log("locale change"));

    let year = getContext("selected_year");
    let month = getContext("selected_month");

    function moveMonth(n: number) {
        let nextMonthDate: Date = new Date($year, $month + n, 1);
        // setContext("selected_month", nextMonthDate.getMonth());
        // setContext("selected_month", nextMonthDate.getMonth());
        $year = nextMonthDate.getFullYear(); //.update((_) => nextMonthDate.getMonth());
        $month = nextMonthDate.getMonth(); //.update((_) => nextMonthDate.getFullYear());
    }
</script>

<div class="container mx-auto">
    <div class="bg-white rounded-lg shadow overflow-hidden">
        <div class="flex items-center justify-between py-2 px-6">
            <div>
                <span class="text-lg font-bold text-gray-800">
                    {$json("date.month_names")[$month]}
                </span>
                <span class="ml-1 text-lg text-gray-600 font-normal">
                    {$year}</span
                >
            </div>
            <div class="border rounded-lg px-1" style="padding-top: 2px;">
                <button
                    type="button"
                    on:click={(e) => {
                        moveMonth(-1);
                    }}
                    class="leading-none rounded-lg transition ease-in-out duration-100 inline-flex cursor-pointer hover:bg-gray-200 p-1 items-center"
                >
                    <svg
                        class="h-6 w-6 text-gray-500 inline-flex leading-none"
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
                <div class="border-r inline-flex h-6" />
                <button
                    type="button"
                    on:click={(e) => {
                        moveMonth(1);
                    }}
                    class="leading-none rounded-lg transition ease-in-out duration-100 inline-flex items-center cursor-pointer hover:bg-gray-200 p-1"
                >
                    <svg
                        class="h-6 w-6 text-gray-500 inline-flex leading-none"
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
