/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.html", "./views/**/*.templ", "./views/**/*.go"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["light"],
  },
}

