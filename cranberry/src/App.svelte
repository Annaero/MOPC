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

  register("en", () => import("./locales/en.json"));
  register("ru", () => import("./locales/ru.json"));

  init({
    fallbackLocale: "en",
    initialLocale: getLocaleFromNavigator(),
  });

  let today: Date = new Date();
  let showYear: number = today.getFullYear();
  let showMonth: number = today.getMonth();

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
</script>

<main class="py-24 relative overflow-scroll">
  <select bind:value={$locale}>
    {#each $locales as locale}
      <option value={locale}>{locale}</option>
    {/each}
  </select>

  {#if $isLoading}
    Please wait...
  {:else}
    <EventCalendar {showYear} {showMonth} {events} />
  {/if}
</main>

<style lang="postcss">
  :global(html) {
    background-color: theme(colors.gray.100);
  }
</style>
