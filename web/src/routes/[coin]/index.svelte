<script>
import { onMount } from 'svelte'
import { page } from '$app/stores';
import { formatPrice } from '@src/functions/utils.js'
import Seo from '@components/Seo.svelte'
import Meme from '@components/Meme.svelte'
import sparkline from "@fnando/sparkline"
import { normalizeArrayToMin, extractObjectProperyToArray } from '@src/functions/utils.js'
import { timeAgo } from '@src/functions/timeago.js'

let API_PATH
if (import.meta.env.VITE_API_PATH !== undefined) {
  API_PATH = import.meta.env.VITE_API_PATH
}

let coin
let coins
let loadingStatus = 'idle'
let sparklineEl
let chartValues
onMount(() => {

  function getCoins() {
    loadingStatus = 'loading'
    coins = JSON.parse(localStorage.getItem('coins'))
    coin = coins.find(element => element.symbol === $page.params.coin)
    fetch(API_PATH+'/markets')
      .then((response) => response.json())
      .then((response) => {
        loadingStatus = 'success'
        coins = response
        localStorage.setItem('coins', JSON.stringify(coins));
        
        chartValues = extractObjectProperyToArray(coin.prices_7d, 'price')
        sparkline(sparklineEl, normalizeArrayToMin(chartValues))
      })
      .catch((error) => {
        loadingStatus = 'error'
      });
  }
  const refreshInterval = setInterval(getCoins, 10*1000);
  getCoins()
  return () => clearInterval(refreshInterval)

})
</script>

<svelte:head>
  {#if coin}
    <Seo
      title={` ${coin.name} Price and ${coin.name} meme - CryptoMeme.WTF`}
      desc={` ${coin.name} Price and ${coin.name} meme - Crypto memes`}
      img={`/card.jpg`}
    />
  {:else}
    <Seo
      title={`${$page.params.coin} - CryptoMeme.WTF`}
      desc={`${$page.params.coin} - Crypto price memes`}
      img={`/card.jpg`}
    />
  {/if}
</svelte:head>

{#if coin}
  <div class="flex gap-14 flex-col-reverse items-center md:flex-row md:items-start">
    <div class="w-full md:max-w-[20rem] rounded-box flex flex-col gap-8 mx-auto max-w-[85vw]">

      <div class="flex gap-2 items-center w-full">
        <img class="w-12 h-12 rounded" src={coin.image} alt={coin.name} width="48" height="48"/>
        <div>
          <h1 class="font-black text-xl">{coin.name}</h1>
          <div class="font-mono text-xs uppercase opacity-80">{coin.symbol}</div>
        </div>
      </div>

      <div class="p-4 rounded-box bg-base-200">
        <div class="font-mono text-xl font-bold">
          ${formatPrice(coin.current_price)}
        </div>
        <div class="flex items-center gap-1">
          {#if coin.price_change_percentage_24h > 0}
            <svg class="w-4 h-4 fill-current text-success" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M17.92,6.62a1,1,0,0,0-.54-.54A1,1,0,0,0,17,6H7A1,1,0,0,0,7,8h7.59l-8.3,8.29a1,1,0,0,0,0,1.42,1,1,0,0,0,1.42,0L16,9.41V17a1,1,0,0,0,2,0V7A1,1,0,0,0,17.92,6.62Z"/></svg>
            <span class="z-[2] font-mono text-sm text-success tooltip tooltip-bottom tooltip-accent" data-tip="+${Number(Math.abs(coin.price_change_24h))} in 24hr">
              {Number(Math.abs(coin.price_change_percentage_24h)).toFixed(2)}%
            </span>
          {:else if coin.price_change_percentage_24h < 0}
            <svg class="w-4 h-4 fill-current text-error" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M17,6a1,1,0,0,0-1,1v7.59L7.71,6.29A1,1,0,0,0,6.29,7.71L14.59,16H7a1,1,0,0,0,0,2H17a1,1,0,0,0,.38-.08,1,1,0,0,0,.54-.54A1,1,0,0,0,18,17V7A1,1,0,0,0,17,6Z"/></svg>
            <span class="z-[2] font-mono text-sm text-error tooltip tooltip-bottom tooltip-accent" data-tip="-${Number(Math.abs(coin.price_change_24h))} in 24hr">
              {Number(Math.abs(coin.price_change_percentage_24h)).toFixed(2)}%
            </span>
          {/if}
        </div>
      </div>

      <div class="px-4"> 
        <div class="font-bold text-sm flex items-center gap-2 mb-2">7 day price change</div>
        <div class="w-full">
          <svg 
            bind:this={sparklineEl} 
            class="mx-auto max-w-full" 
            width="220" 
            height="100" 
            stroke-width="2" 
            fill="none" 
            stroke="hsla(var({coin.prices_7d[0].price > coin.prices_7d[coin.prices_7d.length - 1].price ? '--er' : '--su'}))"
          >
          </svg>
        </div>
      </div>

      <div class="px-4"> 
        <div class="font-bold text-sm flex items-center gap-2 mb-2">Today price range</div>
        <div class="relative pt-8">
          <div class="tooltip absolute tooltip-open tooltip-primary" data-tip="${formatPrice(coin.current_price)}" style="left:{
            Math.min(
              ((coin.high_24h - coin.current_price)*100)/(coin.high_24h - coin.low_24h)
            , 100)
            }%">
            <div class="w-px"></div>
          </div>
        </div>
        <div class="h-2 p-px flex items-stretch bg-base-content bg-opacity-20 rounded-full">
          <div class="bg-base-content rounded-full" style="width:{
            Math.min(
              ((coin.high_24h - coin.current_price)*100)/(coin.high_24h - coin.low_24h)
            , 100)
            }%"></div>
        </div>
        <div class="flex justify-between py-2">
          <div class="font-mono text-xs text-left">${formatPrice(coin.low_24h)}<br/>Min</div>
          <div class="font-mono text-xs text-right">${formatPrice(coin.high_24h)}<br/>Max</div>
        </div>
      </div>

      <div class="px-4"> 
        <div class="font-bold text-sm flex items-center gap-2 mb-2">All time range</div>
        <div class="relative pt-8">
          <div class="tooltip absolute tooltip-open tooltip-primary" data-tip="${formatPrice(coin.current_price)}" style="left:{
            Math.min(
              100 - Number(Math.abs(coin.ath_change_percentage))
            , 100)
            }%">
            <div class="w-px"></div>
          </div>
        </div>
        <div class="h-2 p-px flex items-stretch bg-base-content bg-opacity-20 rounded-full">
          <div class="bg-base-content rounded-full" style="width:{100 - Number(Math.abs(coin.ath_change_percentage))}%"></div>
        </div>
        <div class="flex justify-between py-2">
          <div class="font-mono text-xs text-left">${formatPrice(coin.atl)}<br/>ATL</div>
          <div class="font-mono text-xs text-right">${formatPrice(coin.ath)}<br/>ATH</div>
        </div>
      </div>


      <div class="px-4"> 
        <div class="font-bold text-sm flex items-center gap-2 mb-2">Market Cap</div>
        <div class="font-mono text-sm">${formatPrice(coin.market_cap)}</div>
      </div>

      <div class="px-4"> 
        <div class="font-bold text-sm flex items-center gap-2 mb-2">Circulating Supply</div>
        <div class="font-mono text-sm">{coin.circulating_supply}</div>
      </div>

      <!-- <div class="px-4"> 
        <div class="font-bold text-sm flex items-center gap-2 mb-2">Last update</div>
        <div class="font-mono text-sm">{timeAgo(coin.last_updated, true)} ago</div>
      </div> -->




    </div>
    <div>
      {#if coin.meme}
        <div class="tooltip tooltip-open tooltip-bottom" data-tip="↑ click to download image">
          <div class="my-3">
            <Meme 
              id={coin.symbol} 
              coin_name={coin.name} 
              meme={coin.meme} 
              meme_caption={coin.meme_caption} 
              creditDate=""
              creditPrice= "{coin.symbol.toUpperCase()} price on {new Date(Date.now()).toLocaleString("en-US",{ year: "numeric", month: "short", day: "2-digit" })}: ${coin.current_price} ({(coin.price_change_24h<0?"▼":"▲") + '$' + Number(Math.abs(coin.price_change_24h))})"
              classes='w-full rounded-box'
              clickToDownload=true
            />
          </div>
        </div>
      {/if}
    </div>
  </div>
{:else if loadingStatus == 'error'}
  not found
{:else if loadingStatus == 'success' && !coin}
  not found
{/if}
