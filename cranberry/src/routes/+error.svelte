<script>
    import { page } from "$app/stores";
</script>

<div
    class="flex flex-grow min-h-full items-center justify-center lg:px-8 bg-base-200 rounded-lg shadow-md max-w-md w-full"
>
    <div class="p-8 text-center">
        <h1 class="mb-4 text-4xl font-bold text-error">{$page.status}</h1>
        <p class="text-error-content mb-7 text-xl">{$page.error.message}</p>
        <a
            href="/"
            class="btn btn-xs sm:btn-sm md:btn-md lg:btn-lg btn-active btn-secondary"
        >
            Go back to Home
        </a>
    </div>
</div>
