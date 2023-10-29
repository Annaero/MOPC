<script lang="ts">
    import "$lib/i18n";
    import "../app.css";
    import Icon from "@iconify/svelte";
    import logo from "$lib/assets/mors_4507322.png";
    import { isLoading, locale, locales } from "svelte-i18n";
    import { browser } from "$app/environment";
    // Import to initialize. Important :)
    import { waitLocale } from "svelte-i18n";
    import type { LayoutLoad } from "./$types";
    import { invalidate } from "$app/navigation";
    import { onMount } from "svelte";

    export let data;

    let { supabase, session } = data;
    $: ({ supabase, session } = data);

    onMount(() => {
        const {
            data: { subscription },
        } = supabase.auth.onAuthStateChange((event, _session) => {
            if (_session?.expires_at !== session?.expires_at) {
                invalidate("supabase:auth");
            }
        });

        return () => subscription.unsubscribe();
    });

    export const load: LayoutLoad = async () => {
        if (browser) {
            locale.set(window.navigator.language);
        }
        await waitLocale();
    };

    const handleSignOut = async () => {
        await supabase.auth.signOut();
    };

    const flag = { "en-GB": "🇬🇧", "ru-RU": "🇷🇺" };
    export const ssr = false;
</script>

<div class="mx-auto bg-neutral flex flex-col place-items-center items-center">
    <div class="navbar bg-neutral sm:w-3/4 md:w-1/2 lg:w-3/5">
        <div class="navbar-start">
            <a class="btn btn-ghost normal-case text-xl" href="/events/calendar"
                ><img src={logo} alt="MOPC" class="w-8" />MOPC</a
            >
        </div>

        <div class="navbar-end">
            <div>
                <select
                    bind:value={$locale}
                    class="select select-ghost select-sm w-full max-w-xs text-3xl"
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
                        {#if session}
                            <Icon
                                icon="carbon:user-avatar"
                                height="2.5rem"
                                width="2.5rem"
                            />
                        {:else}
                            <Icon
                                icon="carbon:user-avatar-filled"
                                height="2.5rem"
                                width="2.5rem"
                            />
                        {/if}
                    </div>
                </label>
                <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
                <ul
                    tabindex="0"
                    class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
                >
                    {#if session}
                        <li class="text-bold">
                            <a href="/auth" on:click={handleSignOut}>Sign Out</a
                            >
                        </li>
                    {:else}
                        <li class="text-bold"><a href="/auth">Sign In</a></li>
                    {/if}
                </ul>
            </div>
        </div>
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
</main>

<footer
    class="footer footer-center p-4 bg-base-300 text-base-content h-60px"
    style="position:fixed;"
>
    <aside>
        <p>Copyright © 2023 - All right reserved by Aleksei Krikunov</p>
    </aside>
</footer>

<style lang="postcss">
</style>
