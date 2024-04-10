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
    import { invalidate, afterNavigate, goto } from "$app/navigation";
    import { onMount } from "svelte";
    import { page } from "$app/stores";
    import { redirect } from "@sveltejs/kit";

    export let data;

    let { supabase, session, profile, avatarUrl } = data;
    $: ({ supabase, session, profile } = data);

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

    // Profiles does not creates automatically after user creation
    afterNavigate(({ from, to }) => {
        if (
            session &&
            !profile &&
            !to.url.pathname.startsWith("/profile/edit")
        ) {
            goto("/profile/edit");
        }
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

    const flag = { "en-GB": "EN", "ru-RU": "RU" };
    // export const ssr = false;
</script>

<div class="mx-auto bg-base-300 flex flex-col place-items-center items-center">
    <div class="navbar bg-base-300 w-full md:w-3/4 xl:w-3/5 2xl:w-1/2">
        <div class="navbar-start">
            <a class="btn btn-ghost normal-case text-xl" href="/events/calendar"
                ><img src={logo} alt="MOPC" class="w-8" />MOPC</a
            >
        </div>
        <div class="navbar-center">
            <label class="flex cursor-pointer gap-2">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="20"
                    height="20"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    ><circle cx="12" cy="12" r="5" /><path
                        d="M12 1v2M12 21v2M4.2 4.2l1.4 1.4M18.4 18.4l1.4 1.4M1 12h2M21 12h2M4.2 19.8l1.4-1.4M18.4 5.6l1.4-1.4"
                    /></svg
                >
                <input
                    type="checkbox"
                    value="dark"
                    class="toggle theme-controller"
                />
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    width="20"
                    height="20"
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    ><path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z"
                    ></path></svg
                >
            </label>
        </div>

        <div class="navbar-end">
            <div>
                <select
                    bind:value={$locale}
                    class="select select-ghost select-sm w-full max-w-xs text-l"
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
                <label
                    tabindex="0"
                    class="btn btn-square rounded-xl hover:border hover:border-accent"
                >
                    <div class="w-15 rounded-xl bg-neutral avatar">
                        {#if profile}
                            <div
                                class="w-15 rounded-xl text-neutral-content bg-neutral"
                            >
                                {#if avatarUrl}
                                    <img alt="avatar" src={avatarUrl} />
                                {:else}
                                    <span class="text-xl uppercase">
                                        {profile.username
                                            ? profile.username.slice(0, 2)
                                            : "🙃"}
                                    </span>
                                {/if}
                            </div>
                        {:else}
                            <Icon
                                icon="carbon:user-avatar-filled"
                                height="3rem"
                                width="3rem"
                            />
                        {/if}
                    </div>
                </label>
                <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
                <div
                    tabindex="0"
                    class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-200 rounded-box w-52 content-end"
                >
                    <ul>
                        {#if session}
                            {#if profile}
                                <label
                                    >Hello, <a
                                        class="link link-accent"
                                        href="/profile/{profile.id}"
                                        >{profile.username}</a
                                    >!</label
                                >
                            {:else}
                                <label>Hello, The Nameless One</label>
                            {/if}
                        {:else}
                            <label> Please Sign In </label>
                        {/if}
                        {#if session}
                            <li class="text-bold">
                                <a href="/auth" on:click={handleSignOut}
                                    >Sign Out</a
                                >
                            </li>
                        {:else}
                            <li class="text-bold">
                                <a href="/auth">Sign In</a>
                            </li>
                        {/if}
                    </ul>
                </div>
            </div>
        </div>
    </div>
</div>

<main class="min-h-full">
    <div
        class="py-7 mx-auto flex flex-col place-items-center items-center justify-center mb-60"
    >
        <div
            class="w-full w-full md:w-3/4 xl:w-3/5 2xl:w-1/2 flex flex-row items-center justify-center"
        >
            {#if $isLoading}
                <span class="loading loading-spinner loading-lg" />
            {:else}
                <slot />
            {/if}
        </div>
    </div>
</main>

<footer class="footer footer-center p-4 bg-base-300 text-base-content h-60px">
    <aside>
        <p>Copyright © 2023 - All right reserved by Aleksei Krikunov</p>
    </aside>
</footer>

<style lang="postcss">
</style>
