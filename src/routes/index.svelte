<script>
import TableItem from '@components/TableItem.svelte'
import { onMount } from 'svelte'

let coins
let loadingStatus = 'loading'
onMount(() => {
  coins = JSON.parse(localStorage.getItem('coins'))
  function getCoins() {
    fetch('http://127.0.0.1:8569/v1/markets')
    .then((response) => response.json())
    .then((response) => {
      loadingStatus = 'success'
      coins = response
      localStorage.setItem('coins', JSON.stringify(coins));
    })
    .catch((error) => {
      loadingStatus = 'error'
    })
  }
  getCoins()

})

</script>

<div class="overflow-x-auto">
  <table class="table w-full">
    <thead>
      <tr>
        <th>
          {#if loadingStatus == 'loading'}
            <svg class="animate-spin h-5 w-5 text-base-content" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
          {:else if loadingStatus == 'error'}
            <span class="flex gap-2 items-center">
              <span class="opacity-20 text-xs normal-case">
                Last update: 
                {new Date(coins[0].last_updated).toLocaleString('en-us', {weekday:'short'})}
                {new Date(coins[0].last_updated).getHours()}:{new Date(coins[0].last_updated).getMinutes()}:{new Date(coins[0].last_updated).getSeconds()}
              </span>
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
      {#if coins !== undefined}
        {#each coins as coin}
          <TableItem {coin} />
        {/each}
      {/if}
    </tbody>
  </table>
</div>