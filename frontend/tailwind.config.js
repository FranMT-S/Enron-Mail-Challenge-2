const plugin = require('tailwindcss/plugin')
/** @type {import('tailwindcss').Config} */
export default {
  darkMode:false,
  content: [],
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    screens:{
      'mobile-sm':"361px",
      'mobile':"460px"
    },
    animationDuration:{
      '500':'500ms',
      '800':'800ms',
      '1000':'1000ms',
      '2000':'2000ms',
      '3000':'3000ms',
      '4000':'4000ms',
      '5000':'5000ms',
    }
    ,
    extend: {},
  },
  plugins: [
    plugin(function({ matchUtilities, theme }) {
      matchUtilities(
        {
          'animation-duration': (value) => ({
            'animation-duration': value
          }),
        },
        { values: theme('animationDuration') }
      )
    })

  ],

}

