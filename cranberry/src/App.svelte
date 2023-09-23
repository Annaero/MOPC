<script lang="ts">
  import EventCalendar from "./lib/EventCalendar/EventCalendar.svelte";
  import {
    register,
    init,
    getLocaleFromNavigator,
    locale,
    locales,
  } from "svelte-i18n";
  import type { MOPCEvent } from "./lib/models/mopcEvent";
  import { getEvents } from "./services/event";

  register("en", () => import("./locales/en.json"));
  register("ru", () => import("./locales/ru.json"));

  init({
    fallbackLocale: "en",
    initialLocale: getLocaleFromNavigator(),
  });

  let today: Date = new Date();
  let yearToShow: number = today.getFullYear();
  let monthToShow: number = today.getMonth();

  let events: MOPCEvent[] = [];

  async function updateEvents(year: number, month: number) {
    const firstayOfMonth = new Date(year, month, 1);
    const lastDayOfMonth = new Date(year, month + 1, 0);
    events = await getEvents(firstayOfMonth, lastDayOfMonth);
    console.log(events);
  }
</script>

<main class="py-24 relative overflow-scroll">
  <select bind:value={$locale}>
    {#each $locales as locale}
      <option value={locale}>{locale}</option>
    {/each}
  </select>

  {#await updateEvents(yearToShow, monthToShow)}
    Loading...
  {:then _value}
    <EventCalendar
      showYear={yearToShow}
      showMonth={monthToShow}
      {events}
      on:selectedDateChanged={async (e) => {
        await updateEvents(e.detail.year, e.detail.month);
      }}
    />
  {/await}
</main>

<style lang="postcss">
  :global(html) {
    background-color: theme(colors.gray.100);
  }
</style>
