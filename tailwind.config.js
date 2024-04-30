/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.js", "./views/**/*.templ", "./views/**/*.go"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["forest"],
  },
}

