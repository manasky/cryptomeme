<script>
import { onMount } from 'svelte'
import CardItem from '@components/CardItem.svelte'
import Meme from '@components/Meme.svelte'
import Seo from '@components/Seo.svelte'
import { sparkline } from "@fnando/sparkline"
import { normalizeArrayToMin, extractObjectProperyToArray } from '@src/functions/utils.js'
import { page } from '$app/stores';
import { goto } from '$app/navigation';

// handle dynamic routes
// onMount(() => {
//   if (window.location.pathname != '/') {
//     goto(window.location.pathname) 
//   }
// })

let API_PATH
if (import.meta.env.VITE_API_PATH !== undefined) {
  API_PATH = import.meta.env.VITE_API_PATH
}

let sparklineEl = []
let sparklineEl2 = []

let coins
let lastUpdate
let loadingStatus = 'idle'
onMount(() => {

  function getCoins() {
    loadingStatus = 'loading'
    coins = JSON.parse(localStorage.getItem('coins'))
    lastUpdate = localStorage.getItem('lastUpdate')
    fetch(API_PATH+'/markets')
      .then((response) => response.json())
      .then((response) => {
        loadingStatus = 'success'
        coins = response
        localStorage.setItem('coins', JSON.stringify(coins));
        localStorage.setItem('lastUpdate', new Date().toISOString());
        for (let coin of coins) {
          if(coin.prices_7d != null){
            let chartValues = extractObjectProperyToArray(coin.prices_7d, 'price')
            sparkline(sparklineEl[coin.symbol], normalizeArrayToMin(chartValues))
            sparkline(sparklineEl2[coin.symbol], normalizeArrayToMin(chartValues))
          }
        }
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
  <Seo
    title={`CryptoMeme.WTF - Crypto Price + Memes`}
    desc={`Crypto price memes - Bitcoin meme, Ethereum meme, crypto meme`}
    img={`/card.jpg`}
  />
</svelte:head>

<div class="-mx-3 md:hidden">
  {#if coins}
  <div class="w-full">
    <div class="w-full overflow-y-hidden carousel carousel-center">
      <div class="w-1/12 carousel-item"></div>
      {#each coins.slice(0, 100) as coin}
        <div class="w-10/12 carousel-item">
          <CardItem {coin}>
            {#if coin.meme}
              <div class="cursor-pointer">
                <Meme 
                  coin={coin}
                  classes="w-full my-0"
                  wrapperClasses="rounded-box"
                />
              </div>
            {/if}
            <span class="inline-block mx-auto" slot="weeklyChart">
              {#if coin.prices_7d != null}
                <div class="tooltip" data-tip="7 day price changes">
                  <svg 
                    bind:this={sparklineEl[coin.symbol]} 
                    class="max-w-full mx-auto" 
                    width="250" 
                    height="60" 
                    stroke-width="2" 
                    fill="none" 
                    stroke="hsla(var({coin.prices_7d[0].price > coin.current_price ? '--er' : '--su'}))"
                  >
                  </svg>
                </div>
              {/if}
            </span>
          </CardItem>
        </div>
      {/each}
      <div class="w-1/12 carousel-item"></div>
    </div>
  </div>
  {/if}
</div>


<div class="hidden -mx-3 md:block">
  {#if coins}
    <div class="w-full md:flex md:flex-col">
      {#each coins.slice(0, 100) as coin}
        <CardItem {coin}>
          {#if coin.meme}
            <div class="cursor-pointer">
              <Meme 
                coin={coin}
                classes="w-full my-0 hover:brightness-90"
                wrapperClasses="rounded-xl"
              />
            </div>
          {/if}
            <span slot="weeklyChart">
              {#if coin.prices_7d != null}
                <div class="tooltip" data-tip="7 day price changes">
                  <svg 
                    bind:this={sparklineEl2[coin.symbol]} 
                    class="max-w-full mx-auto" 
                    width="220" 
                    height="60" 
                    stroke-width="2" 
                    fill="none" 
                    stroke="hsla(var({coin.prices_7d[0].price > coin.current_price ? '--er' : '--su'}))"
                  >
                  </svg>
                </div>
              {/if}
            </span>
        </CardItem>
      {/each}
    </div>
  {/if}
</div>