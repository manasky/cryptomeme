<script>
import { onMount } from 'svelte'
import { timeAgo } from '@src/functions/timeAgo.js'
import TableItem from '@components/TableItem.svelte'

let API_PATH
if (import.meta.env.VITE_API_PATH !== undefined) {
  API_PATH = import.meta.env.VITE_API_PATH
}

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

<div class="overflow-x-auto">
  <table class="table w-full">
    <thead>
      <tr>
        <th>
          {#if loadingStatus == 'loading'}
            <svg class="w-4 h-4 animate-spin text-base-content" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
          {:else if loadingStatus == 'error'}
            <span class="text-xs normal-case text-warning">
              {#if lastUpdate}
                Last update: {timeAgo(lastUpdate)}
              {:else}
                Cannot connect to server
              {/if}
            </span>
          {/if}
        </th> 
        <th>Price</th> 
        <th>24h %</th> 
        <th>7d %</th>
        <th>7d chart</th>
      </tr>
    </thead> 
    <tbody>
      {#if coins}
        {#each coins as coin}
          <TableItem {coin} />
        {/each}
      {:else}
        {#each Array(10) as coin}
          <TableItem />
        {/each}
      {/if}
    </tbody>
  </table>
</div>