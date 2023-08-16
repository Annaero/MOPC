<script lang="ts">
    import MonthViewHeader from "./MonthViewHeader.svelte";
    import { locale, json } from "svelte-i18n";
    import { selected_year, selected_month } from "../stores";
    import type { Day } from "../day";
    import type { CalendarEvent } from "../calendarEvent";
    import { isToday } from "../../dateUtils";

    let selected_event = 0;
    let events: CalendarEvent[] = [
        {
            id: 1,
            name: "My fancy event",
            startDate: new Date(2023, 7, 5),
            endDate: new Date(2023, 7, 6),
            active: false,
        },
        {
            id: 2,
            name: "Another event",
            startDate: new Date(2023, 7, 6),
            endDate: new Date(2023, 8, 10),
            active: false,
        },
        {
            id: 3,
            name: "mega cool XXXL",
            startDate: new Date(2023, 7, 23),
            endDate: new Date(2023, 8, 1),
            active: false,
        },
        {
            id: 4,
            name: "small event",
            startDate: new Date(2023, 6, 10),
            endDate: new Date(2023, 7, 3),
            active: false,
        },
    ];

    events.sort((e1, e2) => e1.startDate.getTime() - e2.startDate.getTime());

    locale.set("en");
    locale.subscribe(() => console.log("locale change"));

    // let today: Date = new Date();
    let year: number;
    let month: number;
    let days: Day[] = [];

    selected_year.subscribe((value) => {
        year = value;
        // days = updateCalendar();
    });

    selected_month.subscribe((value) => {
        month = value;
        // days = updateCalendar();
    });

    var firstDayOfMonth: Date;
    var lastDayOfMonth: Date;

    $: {
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

        days = [];
        // why `let` outside the forEach: https://github.com/sveltejs/svelte/issues/6706
        let d = new Date(firstWeekDay);
        events.forEach((e) => (e.active = e.id == selected_event));
        for (; d <= lastWeekDay; d.setDate(d.getDate() + 1)) {
            let day: Day = {
                date: new Date(d),
                today: isToday(d),
                events: events.filter(
                    (v) => v.startDate <= d && v.endDate >= d
                ),
            };
            console.log(day);
            days.push(day);
        }

        // return days;
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
            <div class="flex flex-wrap border-t border-l bg-white">
                {#each days as day}
                    <div
                        style="width: 14.28%; height: 160px"
                        class="pt-2 border-r border-b relative"
                        class:sept3shuf={day.date.getMonth() == 8 &&
                            day.date.getDate() == 3}
                        class:bg-gray-200={day.date.getDay() == 6 ||
                            day.date.getDay() == 0}
                    >
                        <div
                            class="w-6 h-6 mx-4 float-r items-center justify-center cursor-pointer text-center leading-none rounded-full transition ease-in-out duration-100"
                            class:bg-red-400={day.today}
                            class:text-zinc-400={day.date < firstDayOfMonth ||
                                day.date > lastDayOfMonth}
                        >
                            <span>{day.date.getDate()}</span>
                        </div>
                        <ul class="block mt-2 -ml-0.5">
                            {#each day.events as event}
                                {@const firstEventDay =
                                    event.startDate.getTime() ==
                                    day.date.getTime()}
                                {@const lastEventDay =
                                    event.endDate.getTime() ==
                                    day.date.getTime()}

                                <li
                                    class="border-y-2 text-black block overflow-hidden text-left text-xl font-sans my-0.5"
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
                                    {#if event.startDate.getTime() == day.date.getTime() || day.date.getDay() == 1}
                                        {event.name}
                                    {:else}
                                        &nbsp;
                                    {/if}
                                </li>
                            {/each}
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
        background-position: center;
    }
</style>
