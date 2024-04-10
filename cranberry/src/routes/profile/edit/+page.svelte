<script lang="ts">
    import { locale } from "svelte-i18n";
    import type { PageData } from "../[[id]]/$types";
    import { superForm } from "sveltekit-superforms/client";
    import SuperDebug from "sveltekit-superforms/client/SuperDebug.svelte";
    import Icon from "@iconify/svelte";
    import { v4 as uuidv4 } from "uuid";

    locale.set("en");
    locale.subscribe(() => console.log("locale change"));

    export let data: PageData;
    let { profile, supabase } = data;

    const { form, errors, constraints, enhance, delayed, message } = superForm(
        data.form,
        { customValidity: true },
    );

    let avatarUrl;
    let files: FileList;
    let avatarUploading = false;
    let oldAvatar = null;

    const downloadImage = async (path: string) => {
        try {
            const { data, error } = await supabase.storage
                .from("avatars")
                .download(path);

            if (error) {
                throw error;
            }

            avatarUrl = URL.createObjectURL(data);
        } catch (error) {
            if (error instanceof Error) {
                console.error("Error downloading image: ", error.message);
            }
        }
    };

    if ($form.avatarUUID) {
        downloadImage($form.avatarUUID);
    }

    const uploadAvatar = async () => {
        try {
            avatarUploading = true;
            if (!files || files.length === 0) {
                throw new Error("You must select an image to upload.");
            }

            const image = files[0];

            if (image.size > 2097152) {
                throw new Error("File size can't be larger than 2MB");
            }

            const image_name = image.name;
            const uuid = uuidv4();

            const storage_imagename = uuid + "_" + image_name;

            const { error } = await supabase.storage
                .from("avatars")
                .upload(storage_imagename, image);

            if (error) {
                throw error;
            }

            oldAvatar = $form.avatarUUID;
            $form.avatarUUID = storage_imagename;

            await downloadImage(storage_imagename);
            $message = null;
        } catch (error) {
            if (error instanceof Error) {
                console.error("Error uploading avatar:", error);
                $message = error;
            }

            $form.avatarUUID = oldAvatar;
            avatarUploading = false;
        } finally {
            avatarUploading = false;
        }
    };
</script>

<div class="w-full lg:w-3/4 bg-base-200 rounded-lg shadow-md">
    {#if $message}
        <div class="alert alert-error max-w-full mx-8 mt-5">
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

    <div class="card card-side">
        <div class="mx-auto pl-8 pt-8">
            <div class="avatar mx-auto placeholder">
                <div
                    class="w-32 rounded-xl static text-neutral-content"
                    class:bg-neutral={!$form.avatarUUID}
                >
                    {#if avatarUploading}
                        <span
                            class="loading text-xl loading-spinner text-primary"
                        ></span>
                    {:else}
                        <label
                            class="badge badge-m absolute bottom-2 cursor-pointer"
                            for="avatar">Edit avatar</label
                        >
                        {#if avatarUrl}
                            <img src={avatarUrl} alt="Avatar" />
                        {:else}
                            <span class="text-xl uppercase">
                                {profile.username
                                    ? profile.username.slice(0, 2)
                                    : "🙃"}
                            </span>
                        {/if}
                    {/if}
                    <input
                        hidden
                        id="avatar"
                        name="avatar"
                        type="file"
                        accept="image/png, image/jpeg"
                        bind:files
                        on:change={uploadAvatar}
                        disabled={avatarUploading}
                    />
                </div>
            </div>
        </div>
        <div class="card-body">
            <form method="POST" class="w-full" use:enhance>
                <input hidden id="id" name="id" bind:value={$form.id} />
                <input
                    hidden
                    id="userId"
                    name="userId"
                    bind:value={$form.userId}
                />
                <input
                    hidden
                    id="avatarUUID"
                    name="avatarUUID"
                    bind:value={$form.avatarUUID}
                />
                <div class="form-control">
                    {#if profile}
                        <h3 class="card-title">{profile.username}</h3>
                        <input
                            hidden
                            name="username"
                            bind:value={$form.username}
                        />
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
</div>
