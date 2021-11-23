import adapter from '@sveltejs/adapter-static';
import preprocess from 'svelte-preprocess';
import path from 'path'

import { coins } from './src/coins.js'
let coinPathes = coins.map(el => {
  return '/' + el.toLowerCase()
}).slice(0,200)

export default {

  preprocess: [
    preprocess({
      postcss: true
    })
  ],

  kit: {
    prerender: {
			entries: coinPathes,
		},
		adapter: adapter({
			pages: 'public',
			assets: 'public',
			fallback: null
		}),
    vite: {
      resolve: {
        alias: {
          '@src': path.resolve('./src'),
          '@static': path.resolve('./static'),
          '@components': path.resolve('./src/components'),
        }
      }
    }

  }

}
