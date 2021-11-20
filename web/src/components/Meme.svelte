<script>
export let id
export let meme
export let meme_caption
import { onMount } from 'svelte'
import { bgBox, bottomImage, topText } from '@src/functions/canvas.js'
// import { download } from '@src/functions/downloadImage.js'

let canvasEl
let download
onMount(() => {

  let ctx = canvasEl.getContext('2d')


  
  bgBox({
    canvas: canvasEl,
    ctx: ctx,
  });

  topText({
    canvas: canvasEl,
    ctx: ctx,
    text: meme_caption,
  });
  bottomImage({
    canvas: canvasEl,
    ctx: ctx,
    path : meme,
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
  class="w-52 border border-base-content border-opacity-5 cursor-pointer rounded" width="700" height="600" 
  id={id} on:click={download()} bind:this={canvasEl}
></canvas>


