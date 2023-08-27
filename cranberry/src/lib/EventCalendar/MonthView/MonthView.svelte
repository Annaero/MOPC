<script lang="ts">
    import MonthViewHeader from "./MonthViewHeader.svelte";
    import { locale, json } from "svelte-i18n";
    import type { Day, Week } from "../day";
    import type { CalendarEvent } from "../calendarEvent";
    import { isSameDate } from "../../dateUtils";
    import { getContext } from "svelte";

    let selected_event = getContext("selected_event");
    export let events: CalendarEvent[];

    let today: Date = new Date();
    let year = getContext("selected_year");
    let month = getContext("selected_month");
    let weeks: Week[] = [];

    var firstDayOfMonth: Date;
    var lastDayOfMonth: Date;

    $: {
        console.log("Update state");
        firstDayOfMonth = new Date($year, $month, 1);
        let firstWeekDay: Date = new Date($year, $month, 1);
        firstWeekDay.setDate(2 - firstWeekDay.getDay()); //find Monday date

        lastDayOfMonth = new Date($year, $month + 1, 0); // i hate js, why not just Date(..., -1)?
        let weekDay_lastDayOfMonth: number = lastDayOfMonth.getDay();

        let lastWeekDay: Date = new Date(
            $year,
            $month + 1,
            7 - weekDay_lastDayOfMonth
        );
        lastWeekDay = lastWeekDay;

        weeks = [];
        // why `let` outside the forEach: https://github.com/sveltejs/svelte/issues/6706
        let d = new Date(firstWeekDay);
        events.forEach((e) => (e.active = e.id == selected_event));
        console.log(events);
        for (; d <= lastWeekDay; d.setDate(d.getDate() + 7)) {
            const weekStartDate = new Date(d);
            const weekEndDate = new Date(d);
            weekEndDate.setDate(weekEndDate.getDate() + 7);
            let week: Week = {
                startDate: weekStartDate,
                events: events.filter(
                    (event) =>
                        event.startDate <= weekEndDate &&
                        event.endDate >= weekStartDate
                ),
            };
            weeks.push(week);
        }

        // days = [];
        // // why `let` outside the forEach: https://github.com/sveltejs/svelte/issues/6706
        // let d = new Date(firstWeekDay);
        // events.forEach((e) => (e.active = e.id == selected_event));
        // console.log(events);
        // for (; d <= lastWeekDay; d.setDate(d.getDate() + 1)) {
        //     let day: Day = {
        //         date: new Date(d),
        //         today: isSameDate(d, today),
        //         events: events.filter(
        //             (v) => v.startDate <= d && v.endDate >= d
        //         ),
        //     };
        //     days.push(day);
        // }
    }
</script>

<div class="antialiased sans-serif bg-gray-100 h-screen">
    <MonthViewHeader />
    <div class="container mx-auto px-4 py-2 md:px-24">
        <div class="-mx-1 -mb-1">
            <!-- Day names header -->
            <div class="flex flex-wrap">
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
            <div class="flex flex-wrap border-t border-l bg-white">
                {#each weeks as week}
                    <!-- {@const days = Array(7).map((d) => {
                        const ds = new Date(
                            week.startDate.getFullYear(),
                            week.startDate.getMonth(),
                            1
                        );
                        return ds;
                    })}
                    {console.log(days)} -->
                    {#each [0, 1, 2, 3, 4, 5, 6] as dayOfWeek}
                        {@const day = new Date(
                            week.startDate.getFullYear(),
                            week.startDate.getMonth(),
                            week.startDate.getDate() + dayOfWeek
                        )}
                        <div
                            style="width: 14.28%; height: 160px"
                            class="pt-2 border-r border-b relative"
                            class:bg-gray-200={day.getDay() == 6 ||
                                day.getDay() == 0s}
                        >
                            <div
                                class="w-6 h-6 mx-4 float-r items-center justify-center text-center leading-none rounded-full"
                                class:bg-red-400={isSameDate(day, today)}
                                class:text-zinc-400={day < firstDayOfMonth ||
                                    day > lastDayOfMonth}
                            >
                                <span>{day.getDate()}</span>
                            </div>
                        </div>
                    {/each}
                    <!-- <ul class="block mt-2 -ml-0.5">
                            {#each day.events as event}
                                {@const firstEventDay =
                                    event.startDate.getTime() ==
                                    day.date.getTime()}
                                {@const lastEventDay =
                                    event.endDate.getTime() ==
                                    day.date.getTime()}
                                <li
                                    class="border-y-2 relative text-black block text-left text-xl font-sans my-0.5 text-ellipsis"
                                    class:bg-indigo-100={!(
                                        firstEventDay || lastEventDay
                                    )}
                                    class:bg-indigo-300={firstEventDay ||
                                        lastEventDay}
                                    class:border-double={event.active}
                                    class:border-indigo-500={event.active}
                                    class:rounded-l-lg={firstEventDay}
                                    class:rounded-r-lg={lastEventDay}
                                    on:mouseenter={() => {
                                        selected_event = event.id;
                                    }}
                                    on:mouseleave={() => {
                                        selected_event = -1;
                                    }}
                                >
                                    {#if isSameDate(event.startDate, day.date) || day.date.getDay() == 1}
                                        <p
                                            class="pl-2 line-clamp-1 sticky"
                                            class:longname={!lastEventDay &&
                                                day.date.getDay() < 7}
                                        >
                                            {event.name}
                                        </p>
                                    {:else}
                                        &nbsp;
                                    {/if}
                                </li> -->
                    <!-- {/each}
                        </ul>
                    </div> -->
                {/each}
            </div>
        </div>
    </div>
</div>

<style>
    .longname {
        width: 200%;
        z-index: 100;
    }
</style>
