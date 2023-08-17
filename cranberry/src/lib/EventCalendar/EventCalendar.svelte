<script lang="ts">
    import MonthView from "./MonthView/MonthView.svelte";
    import { setContext } from "svelte";
    import { writable } from "svelte/store";

    export let initial_year: number;
    export let initial_month: number;

    let selected_event = writable(-1);
    let selected_year = writable(initial_year);
    let selected_month = writable(initial_month);

    setContext("selected_event", selected_event);
    setContext("selected_year", selected_year);
    setContext("selected_month", selected_month);

    let events = [
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
            name: "mega cool XXXL so loooong",
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

    setContext("events", events);
</script>

<MonthView />
