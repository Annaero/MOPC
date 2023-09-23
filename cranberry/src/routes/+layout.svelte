<script>
    import "../app.css";
    export const ssr = false;
</script>

<div
    class="antialiased sans-serif bg-gray-100 h-screen font-mono mx-auto h-screen flex items-center justify-center"
>
    <div class="sm:w-3/4 md:w-1/2 lg:w-3/5">
        <slot />
    </div>
</div>

<style lang="postcss">
    :global(html) {
        background-color: theme(colors.gray.100);
    }
</style>
