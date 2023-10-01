<script lang="ts">
    import MonthViewHeader from "./MonthViewHeader.svelte";
    import ListView from "./ListView.svelte";
    import { locale, json } from "svelte-i18n";
    import type { Day, Week } from "../day";
    import type { MOPCEvent } from "$lib/models/mopcEvent";
    import { isSameDate } from "$lib/dateUtils";
    import { getContext } from "svelte";

    const DAY_RELATIVE_WIDTH = 14.28; // 100 / 7

    let selected_event = getContext("selected_event");
    export let events: Array<MOPCEvent>;

    let today: Date = new Date();
    let year = getContext("selected_year");
    let month = getContext("selected_month");
    let weeks: Week[] = [];

    var firstDayOfMonth: Date;
    var lastDayOfMonth: Date;

    function getEventLengh(
        eventStartDayOfWeek: number,
        event: MOPCEvent,
        week: Week
    ): number {
        if (event.endDate == null) return 1;
        return event.endDate <= week.endDate
            ? event.endDate.getDay() - eventStartDayOfWeek + 1
            : -1;
    }

    $: {
        console.log("Update state");
        firstDayOfMonth = new Date($year, $month, 1);
        let firstDayToShow: Date = new Date($year, $month, 1);
        firstDayToShow.setDate(2 - firstDayToShow.getDay()); //find Monday date

        lastDayOfMonth = new Date($year, $month + 1, 0); // I hate js dates, why not just Date(..., -1)?

        let lastDayToShow: Date = new Date(
            $year,
            $month + 1,
            7 - lastDayOfMonth.getDay()
        );

        weeks = [];
        // why `let` outside the forEach: https://github.com/sveltejs/svelte/issues/6706
        let d = new Date(firstDayToShow);
        events.forEach((e) => (e.active = e.id == selected_event));
        console.log(events);
        for (; d <= lastDayToShow; d.setDate(d.getDate() + 7)) {
            const weekStartDate = new Date(d);
            const weekEndDate = new Date(d);
            weekEndDate.setDate(weekEndDate.getDate() + 6);
            let week: Week = {
                startDate: weekStartDate,
                endDate: weekEndDate,
                events: events.filter(
                    (event) =>
                        event.startDate <= weekEndDate &&
                        (event.endDate != null
                            ? event.endDate >= weekStartDate
                            : event.startDate >= weekStartDate)
                ),
            };
            weeks.push(week);
        }
    }
</script>

<MonthViewHeader />
<div class="container mx-auto py-2 flex flex row">
    <div class="w-3/12">
        <div
            class="px-2 py-2 text-gray-600 text-sm uppercase tracking-wide font-bold text-left"
        >
            {$json("headers.events")}
        </div>
        <ListView {events} />
    </div>
    <div class="-mx-1 -mb-1 w-9/12">
        <!-- Day names header -->
        <div class="flex flex-row">
            {#each $json("date.weekdays_names") as weekday}
                <div style="width: {DAY_RELATIVE_WIDTH}%" class="px-2 py-2">
                    <div
                        class="text-gray-600 text-sm uppercase tracking-wide font-bold text-center"
                    >
                        {weekday}
                    </div>
                </div>{/each}
        </div>

        <!-- Actual days -->
        <div class="flex flex-col border-t border-l bg-white">
            {#each weeks as week}
                <div class="relative">
                    <!-- Day cells -->
                    <div class="flex flex-row">
                        {#each Array(7).keys() as dayOfWeek}
                            {@const day = new Date(
                                week.startDate.getFullYear(),
                                week.startDate.getMonth(),
                                week.startDate.getDate() + dayOfWeek
                            )}
                            <div
                                style="width: {DAY_RELATIVE_WIDTH}%; height: 160px"
                                class="pt-2 border-r border-b"
                                class:bg-gray-100={day.getDay() == 6 ||
                                    day.getDay() == 0}
                            >
                                <div
                                    class="w-6 h-6 mx-4 -mt-1 text-center rounded-full"
                                    class:bg-red-400={isSameDate(day, today)}
                                    class:text-zinc-400={day <
                                        firstDayOfMonth || day > lastDayOfMonth}
                                >
                                    {day.getDate()}
                                </div>
                            </div>
                        {/each}
                    </div>
                    <!-- Events list -->
                    <ul class="absolute mt-10 z-10 top-0 w-full">
                        {#each week.events as event}
                            {@const eventStartDayOfWeek =
                                event.startDate >= week.startDate
                                    ? event.startDate.getDay() == 0
                                        ? 6
                                        : event.startDate.getDay() - 1
                                    : 0}
                            {@const eventLenthDays = getEventLengh(
                                eventStartDayOfWeek,
                                event,
                                week
                            )}
                            <li
                                class="px-4 py-0.5 mt-0.5 bg-red-100 rounded-lg shadow-lg sm:block cursor-pointer border-black border-solid"
                                style="margin-left: {DAY_RELATIVE_WIDTH *
                                    eventStartDayOfWeek}%; width: {DAY_RELATIVE_WIDTH *
                                    eventLenthDays}%"
                                class:border-double={event.active}
                                class:border-white={event.active}
                                on:mouseenter={() => {
                                    selected_event = event.id;
                                }}
                                on:mouseleave={() => {
                                    selected_event = -1;
                                }}
                            >
                                <p
                                    class="pl-2 line-clamp-1 sticky cursor-pointer"
                                >
                                    <a href="/events/{event.id}">
                                        {event.name}
                                    </a>
                                </p>
                            </li>
                        {/each}
                    </ul>
                </div>
            {/each}
        </div>
    </div>
</div>

<style>
    .longname {
        width: 200%;
        z-index: 100;
    }
</style>
