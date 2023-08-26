<script lang="ts">
  import EventCalendar from "./lib/EventCalendar/EventCalendar.svelte";
  import {
    register,
    init,
    getLocaleFromNavigator,
    locale,
    locales,
    isLoading,
  } from "svelte-i18n";
  import { dateToISODateStr } from "./lib/dateUtils";
  import type { CalendarEvent } from "./lib/EventCalendar/calendarEvent";
  import { onMount } from "svelte";
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

  let events: CalendarEvent[] = [];

  async function updateEvents(year, month) {
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

  {#if $isLoading}{:else}
    <EventCalendar
      showYear={yearToShow}
      showMonth={monthToShow}
      {events}
      on:selectedDateChanged={async (e) => {
        console.log(e);
        await updateEvents(e.detail.year, e.detail.month);
      }}
    />
  {/if}
</main>

<style lang="postcss">
  :global(html) {
    background-color: theme(colors.gray.100);
  }
</style>
