<script>
export let id
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
    text: "When you pay $120 gas fee for a $4 exchange$4 exchange$4 exchange",
  });
  bottomImage({
    canvas: canvasEl,
    ctx: ctx,
    path : "/hide-your-pain.jpg",
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


