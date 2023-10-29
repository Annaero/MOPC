<script lang="ts">
    import { goto } from "$app/navigation";
    import logo from "$lib/assets/mors_4507322.png";
    import Icon from "@iconify/svelte";
    import { boolean } from "zod";

    enum AuthAction {
        auth = "auth",
        register = "register",
        forgotPassword = "forgotPassword",
    }

    const actionLabels = {
        auth: "Sign in",
        register: "Register",
        forgotPassword: "Restore",
    };

    let errorMessage: string = null;
    let infoMessage: string = null;
    let isLoading: boolean = false;
    let authAction: AuthAction = AuthAction.auth;

    const handleAuthAction = async () => {
        switch (authAction) {
            case AuthAction.auth: {
                handleSignIn();
                break;
            }
            case AuthAction.register: {
                handleRegister();
                break;
            }
            case AuthAction.forgotPassword: {
                handleRestorePassword();
                break;
            }
        }
    };

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

    const handleRegister = async () => {
        if (password != confirmPassword) {
            confirmPasswordInput.setCustomValidity("Passwords don't match");
            errorMessage = "Passwords don't match";
            return;
        }
        isLoading = true;
        isLoading = false;

        // if (signInError) {
        //     errorMessage = signInError.message;
        // } else {
        //     goto("/");
        // }
    };

    const handleRestorePassword = async () => {
        isLoading = true;
        const { data, error: resetError } =
            await supabase.auth.resetPasswordForEmail(email, {
                redirectTo: "https://example.com/update-password",
            });
        isLoading = false;

        if (resetError) {
            errorMessage = resetError.message;
        } else {
            infoMessage = "Information was sent to your email.";
            authAction = AuthAction.auth;
        }
    };

    const clearErrors = async () => {
        infoMessage = null;
        errorMessage = null;
        confirmPasswordInput.setCustomValidity("");
    };

    export let data;
    let { supabase } = data;
    $: ({ supabase } = data);

    let email = "";
    let password = "";
    let confirmPassword = "";

    let passwordInput;
    let confirmPasswordInput;
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
        class="mt-5 sm:mx-auto sm:w-full sm:max-w-sm flex flex-col"
        class:blur-sm={isLoading}
    >
        <p
            class="text-error flex flex row items-center place-self-center"
            class:invisible={errorMessage == null}
        >
            <Icon icon="carbon:error" />
            <span class="pl-2">{errorMessage}</span>
        </p>
        <p
            class="text-success flex flex row items-center place-self-center"
            class:invisible={infoMessage == null}
        >
            <Icon icon="carbon:information-filled" />
            <span class="pl-2">{infoMessage}</span>
        </p>
        <form class="space-y-2" on:submit|preventDefault={handleAuthAction}>
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
                        on:input={clearErrors}
                        bind:value={email}
                    />
                    <i />
                </div>
            </div>

            {#if authAction != AuthAction.forgotPassword}
                <div>
                    <div class="flex items-center justify-between">
                        <label for="password" class="label text-sm font-medium"
                            >Password</label
                        >
                        {#if authAction == AuthAction.auth}
                            <div class="text-sm">
                                <button
                                    class="-ml-2 font-semibold text-secondary btn btn-secondary btn-link btn-sm normal-case no-underline"
                                    on:click={() =>
                                        (authAction =
                                            AuthAction.forgotPassword)}
                                    >Forgot password?</button
                                >
                            </div>
                        {/if}
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
                            on:input={clearErrors}
                            bind:this={passwordInput}
                            bind:value={password}
                        />
                    </div>
                </div>
            {/if}
            {#if authAction === AuthAction.register}
                <div>
                    <div class="flex items-center justify-between">
                        <label for="password" class="label text-sm font-medium"
                            >Confirm password</label
                        >
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
                            id="confirmPassword"
                            name="confirmPassword"
                            type="password"
                            autocomplete="current-password"
                            required
                            class="input input-bordered w-full pl-11"
                            on:input={clearErrors}
                            bind:this={confirmPasswordInput}
                            bind:value={confirmPassword}
                        />
                    </div>
                </div>{/if}
            <div class="pt-10">
                <button type="submit" class="btn btn-secondary w-full"
                    >{actionLabels[authAction]}</button
                >
            </div>
        </form>

        <div class="mt-10 text-sm text-normal pb-10">
            {#if authAction != AuthAction.register}
                <p class="text-center">
                    Not a member?
                    <button
                        class="-ml-2 font-semibold text-secondary btn btn-secondary btn-link btn-sm normal-case no-underline"
                        on:click={() => (authAction = AuthAction.register)}
                        >{actionLabels[AuthAction.register]}</button
                    >
                </p>
            {/if}
            {#if authAction != AuthAction.auth}
                <p class="text-center">
                    Already a member?
                    <button
                        class="-ml-2 font-semibold text-secondary btn btn-secondary btn-link btn-sm normal-case no-underline"
                        on:click={() => (authAction = AuthAction.auth)}
                        >{actionLabels[AuthAction.auth]}</button
                    >
                </p>
            {/if}
        </div>
    </div>
</div>
