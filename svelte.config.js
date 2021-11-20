import adapter from '@sveltejs/adapter-static';
import preprocess from 'svelte-preprocess';
import path from 'path'

export default {

  preprocess: [
    preprocess({
      postcss: true
    })
  ],

  kit: {
    
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