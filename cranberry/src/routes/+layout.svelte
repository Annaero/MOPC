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

    const flag = { en: "🇬🇧", ru: "🇷🇺" };

    init({
        fallbackLocale: "en",
        initialLocale: getLocaleFromNavigator(),
    });

    export const ssr = false;
</script>

<div class="navbar bg-base-200">
    <div class="flex-1">
        <a class="btn btn-ghost normal-case text-xl">MOPC</a>
    </div>

    <div>
        <select
            bind:value={$locale}
            class="select select-ghost select-sm w-full max-w-xs"
        >
            {#each $locales as loc}
                <option value={loc} selected={loc === $locale}
                    >{flag[loc]}</option
                >
            {/each}
        </select>
    </div>

    <!-- svelte-ignore a11y-label-has-associated-control -->
    <div class="dropdown dropdown-end">
        <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
        <label tabindex="0" class="btn btn-ghost btn-circle avatar">
            <div class="w-10 rounded-full">
                <!-- svelte-ignore a11y-missing-attribute -->
                <img src="" />
            </div>
        </label>
        <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
        <ul
            tabindex="0"
            class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
        >
            <li><a href="/">Profile</a></li>
            <li><a href="/">Settings</a></li>
            <li><a href="/">Logout</a></li>
        </ul>
    </div>
</div>
<div class="py-7 mx-auto flex flex-col items-center justify-center">
    <div
        class="sm:w-3/4 md:w-1/2 lg:w-3/5 flex flex-row items-center justify-center"
    >
        <slot />
    </div>
</div>

<style lang="postcss">
</style>
