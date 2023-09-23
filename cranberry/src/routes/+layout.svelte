<script>
    import "../app.css";
    export const ssr = false;
</script>

<div class="antialiased sans-serif bg-gray-100 h-screen font-mono mx-48">
    <slot />
</div>

<style lang="postcss">
    :global(html) {
        background-color: theme(colors.gray.100);
    }
</style>
