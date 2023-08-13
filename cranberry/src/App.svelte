<script lang="ts">
  import Calendar from "@event-calendar/core";
  import DayGrid from "@event-calendar/day-grid";
  import List from "@event-calendar/list";

  let events = [
    {
      allDay: true,
      start: "2023-08-12",
      end: "2023-08-14",
      title: "Test event 1",
    },
    {
      allDay: true,
      start: "2023-08-13",
      end: "2023-08-22",
      title: "Test event 2",
    },
  ];

  let list_plugins = [List];
  let list_options = {
    view: "listMonth",
    events: events,
    nowIndicator: true,
    buttonText: {},
  };

  let calendar_plugins = [DayGrid];
  let calendar_options = {
    view: "dayGridMonth",
    events: events,
    nowIndicator: true,
  };
</script>

<main class="py-24 relative overflow-scroll">
  <div class="container mx-auto relative px-4 max-w-6xl">
    <h1 class="text-3xl font-bold underline">Hello world!</h1>

    <div class="flex flex-row flex-wrap">
      <div class="container">
        <Calendar plugins={list_plugins} options={list_options} />
      </div>
      <div class="ml-auto container">
        <Calendar plugins={calendar_plugins} options={calendar_options} />
      </div>
    </div>
  </div>
</main>

<style lang="postcss">
  :global(html) {
    background-color: theme(colors.gray.100);
  }
</style>
