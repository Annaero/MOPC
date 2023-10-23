<script>
    import logo from "$lib/assets/mors_4507322.png";
</script>

<div class="flex min-h-full flex-col justify-center px-6 py-12 lg:px-8">
    <div class="sm:mx-auto sm:w-full sm:max-w-sm">
        <img class="mx-auto h-24 w-auto" src={logo} alt="MOPC" />
        <h2
            class="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-neutral"
        >
            Sign in to your account
        </h2>
    </div>

    <div class="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
        <form class="space-y-3" action="/login" method="POST">
            <div>
                <label for="email" class="label text-sm font-medium"
                    >Email</label
                >
                <div class="mt-2">
                    <input
                        id="email"
                        name="email"
                        type="email"
                        autocomplete="email"
                        required
                        class="input input-bordered w-full"
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
                <div class="mt-2">
                    <input
                        id="password"
                        name="password"
                        type="password"
                        autocomplete="current-password"
                        required
                        class="input input-bordered w-full"
                    />
                </div>
            </div>

            <div>
                <button type="submit" class="btn btn-primary w-full"
                    >Sign in</button
                >
            </div>
        </form>

        <p class="mt-10 text-center text-sm text-gray-500">
            Not a member?
            <a
                href="/register"
                class="font-semibold text-secondary hover:text-secondary-focus"
                >Register</a
            >
        </p>
    </div>
</div>
