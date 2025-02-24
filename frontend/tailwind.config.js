/** @type {import('tailwindcss').Config} */
export default {
  darkMode:false,
  content: [],
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    screens:{
      'mobile':"460px"
    },
    extend: {},
  },
  plugins: [],

}

