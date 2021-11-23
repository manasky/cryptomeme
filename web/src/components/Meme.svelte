<script>
export let coin
export let clickToDownload = false

export let classes = 'w-52'
export let wrapperClasses = '-mt-3 rounded'

import { afterUpdate } from 'svelte'
import { drawMeme } from '@src/functions/canvas.js'
import { formatPrice } from '@src/functions/utils.js'

let canvasEl
let download
afterUpdate(() => {

  let ctx = canvasEl.getContext('2d')
  drawMeme({
    canvas: canvasEl,
    ctx: ctx,
    bgImage : coin.meme,
    caption: coin.meme_caption,
    logo: '/logos/'+coin.symbol+'.png',
    footer_left: {
      title: coin.name,
      desc: new Date(Date.now()).toLocaleString("en-US",{ year: "numeric", month: "short", day: "2-digit" }),
    },
    footer_right: {
      title: '$'+formatPrice(coin.current_price),
      desc: (coin.price_change_24h<0?"▼":"▲") + ' $' + Number(Math.abs(coin.price_change_24h)),
      color : coin.price_change_24h<0?"#ea5333":"#469586",
    },
  });
  download = () => {
    let link = document.createElement('a');
    link.download = coin.name+'-meme-'+new Date(Date.now()).toLocaleString("en-US",{ year: "numeric", month: "short", day: "2-digit" })+' [cryptomeme.wtf].jpg';
    link.href = canvasEl.toDataURL()
    link.click();
  }

})

</script>

<div class="{wrapperClasses} overflow-hidden border border-base-content border-opacity-5 cursor-pointer">
  <canvas
    class="{classes} relative mb-[-12%] -mr-4" width='700' height='700' 
    on:click={clickToDownload ? download() : ''} bind:this={canvasEl}
  ></canvas>
</div>

