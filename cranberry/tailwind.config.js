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
                "primary": "#e25b7a",//"#9d174d",
                "secondary": "#639469", //"#3f6212",
                "accent": "#a6e3a2",//"#115e59",
                "neutral": "#291e00", //"#1a2329",
                "base-100": "#e9e7e7",//"#f9f9f8",//#f3f4f6",//"#e9e7e7", //"#787576",
                "info": "#9bdaed",
                "success": "#3eea9f",
                "warning": "#8e720b",
                "error": "#fc5972",
            }
        },],
    },
};