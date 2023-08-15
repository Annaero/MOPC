<script lang="ts">
    import MonthViewHeader from "./MonthViewHeader.svelte";
    import { locale, json } from "svelte-i18n";
    import { selected_year, selected_month } from "../stores";
    import type { Day } from "../day";
    import { isToday } from "../../dateUtils";

    locale.set("en");
    locale.subscribe(() => console.log("locale change"));

    // let today: Date = new Date();
    let year: number;
    let month: number;
    let days: Day[];

    selected_year.subscribe((value) => {
        year = value;
        days = updateCalendar();
    });

    selected_month.subscribe((value) => {
        month = value;
        days = updateCalendar();
    });

    var firstDayOfMonth: Date;
    var lastDayOfMonth: Date;

    function updateCalendar() {
        firstDayOfMonth = new Date(year, month, 1);
        let firstWeekDay: Date = new Date(year, month, 1);
        firstWeekDay.setDate(2 - firstWeekDay.getDay()); //find Monday date

        lastDayOfMonth = new Date(year, month + 1, 0); // i hate js, why not just Date(..., -1)?
        let weekDay_lastDayOfMonth: number = lastDayOfMonth.getDay();

        let lastWeekDay: Date = new Date(
            year,
            month + 1,
            7 - weekDay_lastDayOfMonth
        );
        lastWeekDay = lastWeekDay;

        let days: Day[] = [];
        // why `let` instead of `var`: https://github.com/sveltejs/svelte/issues/6706
        let d = new Date(firstWeekDay);
        for (; d <= lastWeekDay; d.setDate(d.getDate() + 1)) {
            let day: Day = {
                date: new Date(d),
                today: isToday(d),
                events: [{ name: "some" }],
            };
            days.push(day);
        }

        return days;
    }
</script>

<div class="antialiased sans-serif bg-gray-100 h-screen">
    <MonthViewHeader />
    <div class="container mx-auto px-4 py-2 md:px-24">
        <div class="-mx-1 -mb-1">
            <!-- Day names header -->
            <div class="flex flex-wrap" style="margin-bottom: -40px;">
                {#each $json("date.weekdays_names") as weekday}
                    <div style="width: 14.26%" class="px-2 py-2">
                        <div
                            class="text-gray-600 text-sm uppercase tracking-wide font-bold text-center"
                        >
                            {weekday}
                        </div>
                    </div>{/each}
            </div>

            <!-- Actual days -->
            <div class="flex flex-wrap border-t border-l">
                {#each days as day}
                    <div
                        style="width: 14.28%; height: 160px"
                        class="pt-2 border-r border-b relative"
                        class:sept3shuf={month == 9 && day.date.getDate() == 3}
                    >
                        <!-- add on click and show current day bg-blue-500 text-white': isToday(date) == true, 'text-gray-700 hover:bg-blue-200': isToday(date) == false-->
                        <div
                            class="w-6 h-6 mx-4 float-r items-center justify-center cursor-pointer text-center leading-none rounded-full transition ease-in-out duration-100"
                            class:bg-red-400={day.today}
                            class:text-zinc-400={day.date < firstDayOfMonth ||
                                day.date > lastDayOfMonth}
                        >
                            <span>{day.date.getDate()}</span>
                        </div>
                        <ul class="block mt-2 -ml-2">
                            <li
                                class="bg-indigo-500 text-white block overflow-hidden text-center"
                            >
                                a
                            </li>
                            <li
                                class="bg-indigo-400 text-white block overflow-hidden text-center"
                            >
                                b
                            </li>
                        </ul>
                    </div>
                {/each}
            </div>
        </div>
    </div>
</div>

<style>
    .sept3shuf {
        background-image: url("shuf-3-sep.jpg");
    }
</style>
