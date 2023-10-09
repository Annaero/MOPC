/** @type {import('tailwindcss').Config} */
export default {
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
        extend: {}
    },
    plugins: [require("daisyui")],
    daisyui: {
        themes: ["garden", {
            MOPC: {
                "primary": "#9d174d",
                "secondary": "#3f6212",
                "accent": "#115e59",
                "neutral": "#1a2329",
                "base-100": "#f2f2f0",//#f3f4f6",//"#e9e7e7", //"#787576",
                "info": "#9bdaed",
                "success": "#3eea9f",
                "warning": "#8e720b",
                "error": "#fc5972",
            }
        },],
    },
};