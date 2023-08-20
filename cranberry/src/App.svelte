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
  import type { CalendarEvent } from "./lib/EventCalendar/calendarEvent";
  import { onMount } from "svelte";

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

  async function updateEvents(event) {
    const response = await fetch("/events");
    const events_json = await response.json();
    events = events_json.map((x) => {
      const e: CalendarEvent = {
        id: x.id,
        name: x.name,
        startDate: new Date(Date.parse(x.startDate)),
        endDate: new Date(Date.parse(x.endDate)),
        active: false,
      };
      e.startDate.setHours(0, 0, 0, 0);
      e.endDate.setHours(0, 0, 0, 0);
      return e;
    });
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
        await updateEvents(e);
      }}
    />
  {/if}
</main>

<style lang="postcss">
  :global(html) {
    background-color: theme(colors.gray.100);
  }
</style>
