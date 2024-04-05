<script lang="ts">
    import { locale } from "svelte-i18n";
    import type { PageData } from "../[[id]]/$types";
    import { superForm } from "sveltekit-superforms/client";
    import SuperDebug from "sveltekit-superforms/client/SuperDebug.svelte";
    import Icon from "@iconify/svelte";

    locale.set("en");
    locale.subscribe(() => console.log("locale change"));

    export let data: PageData;

    let { profile } = data;

    const { form, errors, constraints, enhance, delayed, message } = superForm(
        data.form,
        { customValidity: true },
    );
</script>

<div class="card card-side w-3/4 bg-base-200">
    <div class="items-center">
        {#if profile}
            <Icon icon="carbon:image" height="10rem" width="10rem" />
        {:else}
            <Icon icon="carbon:upload" height="10rem" width="10rem" />
        {/if}
    </div>
    <div class="card-body">
        <form method="POST" class="w-full" use:enhance>
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
            <input hidden id="id" name="id" bind:value={$form.id} />
            <input hidden id="userId" name="userId" bind:value={$form.userId} />
            <input
                hidden
                id="avatarUUID"
                name="avatarUUID"
                bind:value={$form.avatarUUID}
            />
            <div class="form-control">
                {#if profile}
                    <h3 class="card-title">{profile.username}</h3>
                    <input hidden name="username" bind:value={$form.username} />
                {:else}
                    <label for="username" class="label">
                        <span class="label-text">Your nickname:</span>
                    </label>
                    <input
                        id="username"
                        name="username"
                        bind:value={$form.username}
                        class="w-full input input-bordered"
                        class:input-error={$errors.username}
                        placeholder="Nickname"
                        {...$constraints.username}
                    />
                {/if}
            </div>
            <div class="form-control">
                <label for="link" class="label">
                    <span class="label-text">Personal page:</span>
                </label>
                <input
                    id="link"
                    name="link"
                    bind:value={$form.link}
                    class="w-full input input-bordered"
                    class:input-error={$errors.link}
                    placeholder="https://my.personal.page/link"
                    {...$constraints.link}
                />
            </div>
            <div class="form-control">
                <label for="about" class="label"
                    ><span class="label-text">About:</span></label
                >
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
            <div class="mt-2 card-actions justify-end">
                <button class="btn bg-primary"> Save </button>
            </div>
        </form>
    </div>
</div>
