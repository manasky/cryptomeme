<script>
import { formatPrice, extractObjectProperyToArray } from '@src/functions/utils.js'
export let coin = {
  name: '',
  symbol: '',
  current_price: '',
  prices_7d: [],
  price_change_24h: '',
  price_change_percentage_24h: '',
  high_24h: '',
  low_24h: '',
  image: 'data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==',
  chart: 'data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==',
}
</script>

<a href="/{coin.symbol}" class="flex flex-col md:grid md:grid-cols-12 gap-3 px-5 w-full md:items-center md:p-2 md:hover:relative rounded-box md:odd:bg-base-200">

  <div class="mx-auto md:order-last md:col-span-3">
    <slot/>
  </div>

  <div class="flex justify-between gap-3 pb-6 md:pb-0 md:col-span-4 md:justify-between md:px-6 md:grid md:grid-cols-2">
    <div class="flex items-center gap-3 cursor-pointer">
      <img class="w-8 h-8 rounded" src={coin.image} alt={coin.name} width="32" height="32"/>
      <div class="flex flex-col pr-4">
        <div class="font-semibold">
          {#if coin.name}
            {coin.name}
          {:else}
            <svg class="w-4 h-4 animate-spin text-base-content" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
          {/if}
        </div>
        <div class="font-mono text-xs text-opacity-70 font-light uppercase text-base-content">{coin.symbol}</div>
      </div>
    </div>
    <div class="flex flex-col items-end md:items-start">
      <div class="font-mono text-lg font-bold">
        {#if coin.current_price}
          ${formatPrice(coin.current_price)}
        {:else}
          <svg class="w-4 h-4 animate-spin text-base-content" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
        {/if}
      </div>
      <div class="flex items-center gap-1 justify-end md:justify-start">
        {#if coin.price_change_percentage_24h > 0}
          <svg class="w-4 h-4 fill-current text-success" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M17.92,6.62a1,1,0,0,0-.54-.54A1,1,0,0,0,17,6H7A1,1,0,0,0,7,8h7.59l-8.3,8.29a1,1,0,0,0,0,1.42,1,1,0,0,0,1.42,0L16,9.41V17a1,1,0,0,0,2,0V7A1,1,0,0,0,17.92,6.62Z"/></svg>
          <span class="z-[2] font-mono text-xs text-success tooltip before:content-[attr(data-tip)] before:text-xs tooltip-left" data-tip="+${Number(Math.abs(coin.price_change_24h))} in 24hr">
            {Number(Math.abs(coin.price_change_percentage_24h)).toFixed(2)}%
          </span>
        {:else if coin.price_change_percentage_24h < 0}
          <svg class="w-4 h-4 fill-current text-error" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M17,6a1,1,0,0,0-1,1v7.59L7.71,6.29A1,1,0,0,0,6.29,7.71L14.59,16H7a1,1,0,0,0,0,2H17a1,1,0,0,0,.38-.08,1,1,0,0,0,.54-.54A1,1,0,0,0,18,17V7A1,1,0,0,0,17,6Z"/></svg>
          <span class="z-[2] font-mono text-xs text-error tooltip before:content-[attr(data-tip)] before:text-xs tooltip-left" data-tip="-${Number(Math.abs(coin.price_change_24h))} in 24hr">
            {Number(Math.abs(coin.price_change_percentage_24h)).toFixed(2)}%
          </span>
        {/if}
      </div>
    </div>

  </div>

  <!-- <div class="flex justify-center gap-4 md:flex-col-reverse md:items-center">
    <div class="flex items-center gap-1 text-xs">
      <svg class="w-2 h-2 fill-current opacity-30" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 2048 2048"><path d="M0 512h2048L1024 1536 0 512z"/></svg>
      <span class="opacity-70">Min: ${formatPrice(Math.min(...extractObjectProperyToArray(coin.prices_7d, 'price')))}</span>
    </div>
    <div class="flex items-center gap-1 text-xs">
      <svg class="w-2 h-2 fill-current opacity-30" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 2048 2048"><path d="M1024 512l1024 1024H0L1024 512z"/></svg>
      <span class="opacity-70">Max: ${formatPrice(Math.max(...extractObjectProperyToArray(coin.prices_7d, 'price')))}</span>
    </div>
  </div> -->

  <div class="flex flex-col md:col-span-2 gap-2 lg:px-4 order-last md:order-2 mx-auto w-full max-w-[20rem]">
    <span class="font-bold text-xs">Performance</span>
    <div class="flex flex-col justify-center items-center gap-px">
      
      <div class="grid grid-cols-5 gap-2 w-full place-items-center justify-items-start">
        <span class="text-xs opacity-80 col-span-2">24&nbsp;hr</span>
        <div class="h-2 w-full p-px col-span-3 flex items-stretch bg-base-content bg-opacity-20 rounded-full">
          <div class="bg-base-content rounded-full" style="width:{
            100 - Math.min(
              ((coin.high_24h - coin.current_price)*100)/(coin.high_24h - coin.low_24h)
            , 100)
            }%"></div>
        </div>
      </div>

      {#if coin.prices_7d != null}
        <div class="grid grid-cols-5 gap-2 w-full place-items-center justify-items-start">
          <span class="text-xs opacity-80 col-span-2">7&nbsp;day</span>
          <div class="h-2 w-full p-px col-span-3 flex items-stretch bg-base-content bg-opacity-20 rounded-full">
            <div class="bg-base-content rounded-full" style="width:{
              100 - Math.min(
                ((Math.max(...extractObjectProperyToArray(coin.prices_7d, 'price')) - coin.current_price)*100)/(Math.max(...extractObjectProperyToArray(coin.prices_7d, 'price')) - Math.min(...extractObjectProperyToArray(coin.prices_7d, 'price')))
              , 100)
              }%"></div>
          </div>
        </div>
      {/if}

      <div class="grid grid-cols-5 gap-2 w-full place-items-center justify-items-start">
        <span class="text-xs opacity-80 col-span-2">All&nbsp;time</span>
        <div class="h-2 w-full p-px col-span-3 flex items-stretch bg-base-content bg-opacity-20 rounded-full">
          <div class="bg-base-content rounded-full" style="width:{
            100 - Math.min(
              ((coin.ath - coin.current_price)*100)/(coin.ath - coin.atl)
            , 100)
            }%"></div>
        </div>
      </div>


    </div>
  </div>

  <div class="md:col-span-3 text-center order-3">
    <slot name="weeklyChart" />
  </div>


</a>
