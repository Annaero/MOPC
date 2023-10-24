<script lang="ts">
    import { goto } from "$app/navigation";
    import logo from "$lib/assets/mors_4507322.png";
    import Icon from "@iconify/svelte";
    import { boolean } from "zod";

    let errorMessage: string = null;
    let isLoading: boolean = false;

    const handleSignIn = async () => {
        isLoading = true;
        const { error: signInError } = await supabase.auth.signInWithPassword({
            email,
            password,
        });
        isLoading = false;

        if (signInError) {
            errorMessage = signInError.message;
        } else {
            goto("/");
        }
    };

    // const handleForgotPassword;

    export let data;
    let { supabase } = data;
    $: ({ supabase } = data);

    let email = "";
    let password = "";
</script>

<div
    class="flex min-h-full flex-col justify-center lg:px-8 bg-base-200 rounded-lg shadow-md max-w-md w-full"
>
    <div class="sm:mx-auto sm:w-full sm:max-w-sm pt-10">
        <img class="mx-auto h-24 w-auto" src={logo} alt="MOPC" />
        <h2
            class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-neutral"
        >
            Sign in to your account
        </h2>
    </div>
    <div
        class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm flex flex-col"
        class:blur-sm={isLoading}
    >
        <p
            class="text-error flex flex row items-center place-self-center"
            class:invisible={errorMessage == null}
        >
            <Icon icon="carbon:error" />
            <span class="pl-2">{errorMessage}</span>
        </p>
        <form class="space-y-2" on:submit|preventDefault={handleSignIn}>
            <div>
                <label for="email" class="label text-sm font-medium"
                    >Email</label
                >
                <div class="mt-1 flex">
                    <span class="absolute ml-1.5 place-self-center">
                        <Icon
                            icon="carbon:email"
                            style="opacity: 20%"
                            height="2rem"
                            width="2rem"
                        />
                    </span>
                    <input
                        id="email"
                        name="email"
                        type="email"
                        autocomplete="email"
                        required
                        class="input input-bordered w-full pl-11"
                        on:input={() => (errorMessage = null)}
                        bind:value={email}
                    />
                </div>
            </div>

            <div>
                <div class="flex items-center justify-between">
                    <label for="password" class="label text-sm font-medium"
                        >Password</label
                    >
                    <div class="text-sm">
                        <a
                            href="#"
                            class="font-semibold text-secondary hover:text-secondary-focus"
                            >Forgot password?</a
                        >
                    </div>
                </div>
                <div class="mt-1 flex">
                    <span class="absolute ml-1.5 place-self-center">
                        <Icon
                            icon="carbon:password"
                            style="opacity: 20%"
                            height="2rem"
                            width="2rem"
                        />
                    </span>
                    <input
                        id="password"
                        name="password"
                        type="password"
                        autocomplete="current-password"
                        required
                        class="input input-bordered w-full pl-11"
                        on:input={() => (errorMessage = null)}
                        bind:value={password}
                    />
                </div>
            </div>

            <div class="pt-10">
                <button type="submit" class="btn btn-secondary w-full"
                    >Sign in</button
                >
            </div>
        </form>

        <p class="mt-10 text-center text-sm text-normal pb-10">
            Not a member?
            <a
                href="/register"
                class="font-semibold text-secondary hover:text-secondary-focus"
                >Register</a
            >
        </p>
    </div>
</div>
