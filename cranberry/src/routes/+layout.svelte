<script>
    import "../app.css";
    import {
        register,
        init,
        getLocaleFromNavigator,
        locale,
        locales,
    } from "svelte-i18n";

    register("en", () => import("$lib/locales/en.json"));
    register("ru", () => import("$lib/locales/ru.json"));

    init({
        fallbackLocale: "en",
        initialLocale: getLocaleFromNavigator(),
    });

    export const ssr = false;
</script>

<div
    class="antialiased sans-serif bg-gray-100 h-screen font-mono mx-auto h-screen flex flex-col items-center justify-center"
>
    <div>
        <select bind:value={$locale}>
            {#each $locales as locale}
                <option value={locale}>{locale}</option>
            {/each}
        </select>
    </div>
    <div class="sm:w-3/4 md:w-1/2 lg:w-3/5">
        <slot />
    </div>
</div>

<style lang="postcss">
    :global(html) {
        background-color: theme(colors.gray.100);
    }
</style>
