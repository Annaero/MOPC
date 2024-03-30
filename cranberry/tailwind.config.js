/** @type {import('tailwindcss').Config} */
export default {
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
        extend: {}
    },
    plugins: [require("@tailwindcss/typography"), require("daisyui")],
    daisyui: {
        themes: ["garden",
            //{
            //     MOPC: {
            //         "primary": "#ec4899",
            //         "secondary": "#14b8a6",
            //         "accent": "#84cc16",
            //         "neutral": "#0d1024", //"#1a2329",
            //         "base-100": "#303457",//"#eeecec", //"#e9e7e7",//"#f9f9f8",//#f3f4f6",//"#e9e7e7", //"#787576",
            //         "info": "#9bdaed",
            //         "success": "#2cc97d",//"#3eea9f",
            //         "warning": "#8e720b",
            //         "error": "#fc5972",
            //     }
            // },
        ],
    },
};