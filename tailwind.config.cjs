module.exports = {
  mode: 'jit',
  purge: {
    content: [
      './src/**/*.{html,js,svelte,ts}',
    ],
  },
  theme: {
    extend: {
      fontFamily: {
        'sans': ['"Nunito Sans"',"sans-serif"],
        'mono': ['"Cousine"',"monospace"],
      }
    }
  },
  plugins: [
    require('daisyui')
  ],
}