<script lang="ts">
    import { dictionary, json, locale } from "svelte-i18n";
    import type { MOPCEvent } from "$lib/models/event.ts_";
    import { Input, Result } from "postcss";
    import { dateToISODateStr } from "$lib/dateUtils";
    import dayjs from "dayjs";

    locale.set("en");
    locale.subscribe(() => console.log("locale change"));

    const format = "YYYY-MM-DD";

    export let event: MOPCEvent;
    export let showInline: boolean = false;
    export let submitAction: () => void;

    let validation_error = {};

    const descriptionError = false;

    function submitClick() {
        if (oneDay) {
            event.endDate = null;
        } else {
            event.endDate = dayjs(endDateVal, format).toDate();
        }
        event.startDate = dayjs(startDateVal, format).toDate();
        const result = eventSchema.safeParse(event);
        if (result.success) {
            submitAction();
            return;
        }
        result.forEach((error) => {});
    }

    function update_event(field: string, value: any) {
        event[field] = value;
    }

    let oneDay: boolean = event.endDate == null;
    let startDateVal = dayjs(event.startDate).format(format);
    let endDateVal =
        event.endDate == null ? null : dayjs(event.endDate).format(format);
</script>

<div class="max-w-md w-full bg-base-100 p-8 rounded-lg shadow-md">
    <form method="Post" class="flex flex-col w-full form-control">
        <div class="mb-6">
            <label for="eventName" class="label">
                <span class="label-text">Event name</span>
                {#if validation_error["name"] != null}
                    <span class="label-text-alt">validation_error["name"]</span>
                {/if}
            </label>
            <input
                bind:value={event.name}
                class="w-full input input-bordered"
                placeholder="Event name"
            />
        </div>

        <!-- Type selector -->
        <div class="mb-6">
            <label for="eventType" class="label">Event type:</label>
            <textarea
                id="eventType"
                name="type"
                rows="4"
                bind:value={event.description}
                class="w-full textarea textarea-bordered"
                class:textarea-error={descriptionError}
                placeholder="Add your event description"
            />
        </div>

        <!-- Description field -->
        <div class="mb-6">
            <label for="eventDescription" class="label">Description:</label>
            <textarea
                id="eventDescription"
                name="eventDescription"
                rows="4"
                bind:value={event.description}
                class="w-full textarea textarea-bordered"
                class:textarea-error={descriptionError}
                placeholder="Add your event description"
            />
        </div>
        <div class="mb-6 grid grid-cols-2 justify-stretch w-full">
            <div class="w-full pr-1">
                <label for="startDate" class="label">Start date:</label>
                <input
                    id="startDate"
                    type="date"
                    class="input input-bordered w-full"
                    bind:value={startDateVal}
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
                    type="date"
                    class="input input-bordered w-full disabled:text-slate-300"
                    bind:value={endDateVal}
                    min={startDateVal}
                    disabled={oneDay}
                />
            </div>
        </div>

        <!-- Submit Button and Character Limit Section -->
        <div class="flex items-end place-content-end">
            <button
                type="submit"
                on:click={submitClick}
                class="btn bg-primary text-primary-content hover:bg-primary-focus"
            >
                Post
            </button>
        </div>
    </form>
</div>
