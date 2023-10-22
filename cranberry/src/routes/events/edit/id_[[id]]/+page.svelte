<script lang="ts">
    export const ssr = false;
    import { dictionary, json, locale } from "svelte-i18n";
    import type { PageData } from "../[[id]]/$types";
    import { dateProxy, superForm } from "sveltekit-superforms/client";
    import SuperDebug from "sveltekit-superforms/client/SuperDebug.svelte";

    locale.set("en");
    locale.subscribe(() => console.log("locale change"));

    export let data: PageData;
    const { form, errors, constraints, enhance, delayed, message } = superForm(
        data.form,
        {
            onSubmit({ formData }) {
                console.log("oneDay", oneDay);
                if (oneDay) {
                    formData.delete("endDate");
                }
            },
        }
    );

    const proxyStartDate = dateProxy(form, "startDate", { format: "date" });
    const proxyEndDate = dateProxy(form, "endDate", { format: "date" });

    let oneDay = $form["endDate"] == null;
</script>

<div class="max-w-md w-full bg-base-100 p-8 rounded-lg shadow-md">
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
        <div class="mb-6">
            <label for="eventName" class="label"> Event name: </label>
            <input
                id="eventName"
                name="name"
                bind:value={$form.name}
                class="w-full input input-bordered"
                class:input-error={$errors.name}
                placeholder="Event name"
                {...$constraints.name}
            />
        </div>
        <!-- Type selector -->
        <div class="mb-6">
            <label for="eventType" class="label">Event type:</label>
            <select
                id="eventType"
                name="type"
                bind:value={$form.type}
                class="select select-bordered w-full max-w-xs"
                class:select-error={$errors.description}
            >
                {#each data.eventTypes as event_type}
                    <option value={event_type}>{event_type}</option>
                {/each}
            </select>
        </div>
        <div class="mb-6">
            <label for="eventDescription" class="label">Description:</label>
            <textarea
                id="eventDescription"
                name="description"
                rows="4"
                bind:value={$form.description}
                class="w-full textarea textarea-bordered"
                placeholder="Add your event description"
                {...$constraints.description}
            />
        </div>
        <div class="mb-6 grid grid-cols-2 justify-stretch w-full">
            <div class="w-full pr-1">
                <label for="startDate" class="label">Start date:</label>
                <input
                    id="startDate"
                    name="startDate"
                    type="date"
                    class="input input-bordered w-full"
                    bind:value={$proxyStartDate}
                    class:input-error={$errors.startDate}
                    {...$constraints.startDate}
                />
                <div class="mt-1 w-full">
                    <label for="oneDay" class="label">
                        <span class="label-text"> One day</span>
                        <input
                            type="checkbox"
                            class="checkbox checkbox-sm"
                            id="oneDay"
                            bind:checked={oneDay}
                        />
                    </label>
                </div>
            </div>
            <div class="w-full pl-1">
                <label for="endDate" class="label" class:text-slate-300={oneDay}
                    >End date:</label
                >
                <input
                    id="endDate"
                    name="endDate"
                    type="date"
                    class="input input-bordered w-full disabled:text-slate-300"
                    bind:value={$proxyEndDate}
                    class:input-error={$errors.endDate}
                    {...$constraints.endDate}
                    min={$proxyStartDate}
                    disabled={oneDay}
                />
            </div>
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
