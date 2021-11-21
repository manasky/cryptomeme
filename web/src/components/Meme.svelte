<script>
export let id
export let meme
export let meme_caption
export let creditPrice = ''
export let creditDate = ''
import { onMount } from 'svelte'
import { drawMeme } from '@src/functions/canvas.js'

let canvasEl
let download
onMount(() => {

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
    link.download = id+'-meme.jpg';
    link.href = canvasEl.toDataURL()
    link.click();
  }

})

</script>

<canvas
  class="w-52 -my-3 -mr-4 border border-base-content border-opacity-5 cursor-pointer rounded" width="700" height="600" 
  id={id} on:click={download()} bind:this={canvasEl}
></canvas>


