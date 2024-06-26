<script lang="ts">
    import { json } from "svelte-i18n";
    import type { Week } from "../day";
    import type { Event } from "$lib/models/event";
    import { isSameDate } from "$lib/dateUtils";
    import { getContext } from "svelte";

    const DAY_RELATIVE_WIDTH = 14.28; // 100 / 7

    let selected_event = getContext("selected_event");
    export let events: Array<Event>;

    let today: Date = new Date();
    let year = getContext("selected_year");
    let month = getContext("selected_month");
    let weeks: Week[] = [];

    var firstDayOfMonth: Date;
    var lastDayOfMonth: Date;

    function getEventLengh(
        eventStartDayOfWeek: number,
        event: Event,
        week: Week
    ): number {
        if (event.endDate == null) return 1;
        return event.endDate <= week.endDate
            ? event.endDate.getDay() - eventStartDayOfWeek + 1
            : -1;
    }

    $: {
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

<div class="py-2 flex flex row justify-center w-full">
    <div class="-mx-1 -mb-1 2xl:w-10/12 xl:w-11/12 lg:w-full rounded-box">
        <!-- Day names header -->
        <div class="flex flex-row">
            {#each $json("date.weekdays_names") as weekday}
                <div style="width: {DAY_RELATIVE_WIDTH}%" class="px-2 py-2">
                    <div
                        class="text-neutral text-sm uppercase tracking-wide font-bold text-center"
                    >
                        {weekday}
                    </div>
                </div>{/each}
        </div>

        <!-- Weeks -->
        <div
            class="flex flex-col bg-base-100 rounded-box border border-base-200 shadow-lg overflow-hidden"
        >
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
                            {@const isToday = isSameDate(day, today)}
                            <div
                                style="width: {DAY_RELATIVE_WIDTH}%"
                                class="pt-2 shadow h-32"
                                class:bg-base-300={day.getDay() == 6 ||
                                    day.getDay() == 0}
                            >
                                <!-- date label  -->
                                <div
                                    class="w-6 h-6 mx-4 -mt-1 text-center rounded-full"
                                    class:bg-primary-focus={isToday}
                                    class:text-secondary-content={isToday}
                                    class:text-base-200={day <
                                        firstDayOfMonth || day > lastDayOfMonth}
                                >
                                    {day.getDate()}
                                </div>
                            </div>
                        {/each}
                    </div>
                    <!-- Events list -->
                    <ul class="absolute mt-10 z-10 top-0 w-full px-2">
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
                                class="py-0.5 mt-0.5 bg-secondary text-secondary-content rounded-box hover:bg-secondary-focus shadow-lg sm:block cursor-pointer border-black border-solid"
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
                                    <a href="/events/id_{event.id}">
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
