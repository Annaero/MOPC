<script lang="ts">
    import "$lib/i18n";
    import "../app.css";
    import { isLoading, locale, locales } from "svelte-i18n";
    import { browser } from "$app/environment";
    // Import to initialize. Important :)
    import { waitLocale } from "svelte-i18n";
    import type { LayoutLoad } from "./$types";

    export const load: LayoutLoad = async () => {
        if (browser) {
            locale.set(window.navigator.language);
        }
        await waitLocale();
    };

    const flag = { "en-GB": "🇬🇧", "ru-RU": "🇷🇺" };
    export const ssr = false;
</script>

<div class="navbar bg-base-200">
    <div class="flex-1">
        <a class="btn btn-ghost normal-case text-xl">MOPC</a>
    </div>

    <div>
        <select
            bind:value={$locale}
            class="select select-ghost select-sm w-full max-w-xs text-xl"
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

<main class="min-h-full">
    <div
        class="py-7 mx-auto flex flex-col place-items-center items-center justify-center mb-60"
    >
        <div
            class="sm:w-3/4 md:w-1/2 lg:w-3/5 flex flex-row items-center justify-center"
        >
            {#if $isLoading}
                <span class="loading loading-spinner loading-lg" />
            {:else}
                <slot />
            {/if}
        </div>
    </div>
    <footer
        class="footer footer-center p-4 bg-base-300 text-base-content"
        style="position:fixed; height:60px;"
    >
        <aside>
            <p>Copyright © 2023 - All right reserved by Aleksei Krikunov</p>
        </aside>
    </footer>
</main>

<style lang="postcss">
</style>
