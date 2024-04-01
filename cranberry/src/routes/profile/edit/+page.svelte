<script lang="ts">
    import { locale } from "svelte-i18n";
    import type { PageData } from "../[[id]]/$types";
    import { superForm } from "sveltekit-superforms/client";
    import SuperDebug from "sveltekit-superforms/client/SuperDebug.svelte";

    locale.set("en");
    locale.subscribe(() => console.log("locale change"));

    export let data: PageData;
    const { form, errors, constraints, enhance, delayed, message } = superForm(
        data.form,
        { customValidity: true },
    );
</script>

<div class="max-w-md w-full bg-base-200 p-8 rounded-lg shadow-md">
    <form method="POST" class="flex flex-col w-full form-control" use:enhance>
        {#if $message}
            <div class="alert alert-error w-full">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="stroke-current shrink-0 h-6 w-6"
                    fill="none"
                    viewBox="0 0 24 24"
                    ><path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z"
                    /></svg
                >
                <p class="max-w-full break-all">{$message}</p>
            </div>
        {/if}
        <input hidden name="id" bind:value={$form.id} />
        <input hidden name="userId" bind:value={$form.userId} />
        <div class="mb-6">
            <label for="username" class="label"> Your nickname: </label>
            <input
                id="username"
                name="username"
                bind:value={$form.username}
                class="w-full input input-bordered"
                class:input-error={$errors.username}
                placeholder="Nickname"
                {...$constraints.username}
            />
        </div>
        <div class="mb-6">
            <label for="link" class="label"> Personal page: </label>
            <input
                id="link"
                name="link"
                bind:value={$form.link}
                class="w-full input input-bordered"
                class:input-error={$errors.link}
                placeholder="link"
                {...$constraints.link}
            />
        </div>
        <div class="mb-6">
            <label for="about" class="label">About:</label>
            <textarea
                id="about"
                name="about"
                rows="4"
                bind:value={$form.about}
                class="w-full textarea textarea-bordered"
                placeholder="Write something about you"
                {...$constraints.about}
            />
        </div>
        <!-- Submit Button and Character Limit Section -->
        <div class="flex items-end place-content-end">
            <button
                class="btn bg-primary text-primary-content
                hover:bg-primary-focus"
            >
                Submit
            </button>
        </div>
    </form>
</div>
