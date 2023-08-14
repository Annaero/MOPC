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
                    <!-- <template x-for="(date, dateIndex) in no_of_days" :key="dateIndex">	 -->
                    <div
                        style="width: 14.28%; height: 160px"
                        class="px-4 pt-2 border-r border-b relative"
                        class:sept3shuf={month == 9 && day.date.getDate() == 3}
                    >
                        <!-- add on click and show current day bg-blue-500 text-white': isToday(date) == true, 'text-gray-700 hover:bg-blue-200': isToday(date) == false-->
                        <div
                            class="inline-flex w-6 h-6 items-center justify-center cursor-pointer text-center leading-none rounded-full transition ease-in-out duration-100"
                            class:bg-red-400={day.today}
                            class:text-zinc-400={day.date < firstDayOfMonth ||
                                day.date > lastDayOfMonth}
                        >
                            {day.date.getDate()}
                        </div>
                        <!-- <div style="height: 80px;" class="overflow-y-auto mt-1"> -->
                        <!-- <div 
                            class="absolute top-0 right-0 mt-2 mr-2 inline-flex items-center justify-center rounded-full text-sm w-6 h-6 bg-gray-700 text-white leading-none"
                            x-show="events.filter(e => e.event_date === new Date(year, month, date).toDateString()).length"
                            x-text="events.filter(e => e.event_date === new Date(year, month, date).toDateString()).length"></div> -->

                        <!-- <template x-for="event in events.filter(e => new Date(e.event_date).toDateString() ===  new Date(year, month, date).toDateString() )">	
                            <div
                                class="px-2 py-1 rounded-lg mt-1 overflow-hidden border"
                                :class="{
                                    'border-blue-200 text-blue-800 bg-blue-100': event.event_theme === 'blue',
                                    'border-red-200 text-red-800 bg-red-100': event.event_theme === 'red',
                                    'border-yellow-200 text-yellow-800 bg-yellow-100': event.event_theme === 'yellow',
                                    'border-green-200 text-green-800 bg-green-100': event.event_theme === 'green',
                                    'border-purple-200 text-purple-800 bg-purple-100': event.event_theme === 'purple'
                                }"
                            >
                                <p x-text="event.event_title" class="text-sm truncate leading-tight"></p>
                            </div>
                        </template> -->
                        <!-- </div> -->
                    </div>
                    <!-- </template> -->
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
