<script>
export let id
export let meme
export let coin_name = ''
export let meme_caption
export let creditPrice = ''
export let creditDate = ''
export let clickToDownload = false

export let classes = 'w-52 rounded'

import { afterUpdate } from 'svelte'
import { drawMeme } from '@src/functions/canvas.js'

let canvasEl
let download
afterUpdate(() => {

  let ctx = canvasEl.getContext('2d')
  drawMeme({
    canvas: canvasEl,
    ctx: ctx,
    path : meme,
    meme_caption: meme_caption,
    creditPrice: creditPrice,
    creditDate: creditDate,
  });
  download = () => {
    let link = document.createElement('a');
    link.download = coin_name+'-meme-'+new Date(Date.now()).toLocaleString("en-US",{ year: "numeric", month: "short", day: "2-digit" })+' [cryptomeme.wtf].jpg';
    link.href = canvasEl.toDataURL()
    link.click();
  }

})

</script>

<canvas
  class="{classes} -my-3 -mr-4 border border-base-content border-opacity-5 cursor-pointer" width='700' height='600' 
  id={id} on:click={clickToDownload ? download() : ''} bind:this={canvasEl}
></canvas>


