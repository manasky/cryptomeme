<script>
  export let coin = {
    name: '',
    symbol: '',
    current_price: '',
    price_change_24h: '',
    price_change_percentage_24h: '',
    high_24h: '',
    low_24h: '',
    image: 'data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==',
    chart: 'data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==',
  }
</script>

<tr class="hover">
  <th class="flex-shrink-0 rounded-l-lg bg-primary">
    <div class="flex items-center gap-3">
      <img class="w-8 h-8 rounded" src={coin.image} alt={coin.name}/>
      <div class="flex flex-col pr-4">
        <div class="font-semibold">
          {#if coin.name}
            {coin.name}
          {:else}
            <svg class="w-4 h-4 animate-spin text-base-content" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
          {/if}
        </div>
        <div class="font-mono text-xs font-bold text-opacity-50 uppercase text-base-content">{coin.symbol}</div>
      </div>
    </div>
  </th>
  <td>
    <div class="flex flex-col">
      <div class="font-mono text-lg">
        {#if coin.current_price}
          ${coin.current_price.toLocaleString()}
        {:else}
          <svg class="w-4 h-4 animate-spin text-base-content" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle><path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path></svg>
        {/if}
      </div>
      <div class="flex items-center gap-1">
        {#if coin.price_change_percentage_24h > 0}
          <svg class="w-4 h-4 fill-current text-success" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M17.92,6.62a1,1,0,0,0-.54-.54A1,1,0,0,0,17,6H7A1,1,0,0,0,7,8h7.59l-8.3,8.29a1,1,0,0,0,0,1.42,1,1,0,0,0,1.42,0L16,9.41V17a1,1,0,0,0,2,0V7A1,1,0,0,0,17.92,6.62Z"/></svg>
          <span class="z-[2] font-mono text-sm text-success tooltip tooltip-right" data-tip="+${Number(Math.abs(coin.price_change_24h)).toFixed(2)} in 24 hours">
            {Number(Math.abs(coin.price_change_percentage_24h)).toFixed(2)}%
          </span>
        {:else if coin.price_change_percentage_24h < 0}
          <svg class="w-4 h-4 fill-current text-error" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M17,6a1,1,0,0,0-1,1v7.59L7.71,6.29A1,1,0,0,0,6.29,7.71L14.59,16H7a1,1,0,0,0,0,2H17a1,1,0,0,0,.38-.08,1,1,0,0,0,.54-.54A1,1,0,0,0,18,17V7A1,1,0,0,0,17,6Z"/></svg>
          <span class="z-[2] font-mono text-sm text-error tooltip tooltip-right" data-tip="-${Number(Math.abs(coin.price_change_24h)).toFixed(2)} in 24 hours">
            {Number(Math.abs(coin.price_change_percentage_24h)).toFixed(2)}%
          </span>
        {/if}
      </div>
    </div>
  </td>
  <td>
    <div class="flex flex-col">
      <div class="flex items-center gap-1 text-xs">
        <svg class="w-2 h-2 fill-current opacity-30" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 2048 2048"><path d="M1024 512l1024 1024H0L1024 512z"/></svg>
        <span class="opacity-70">Max: ${coin.high_24h.toLocaleString()}</span>
      </div>
      <div class="flex items-center gap-1 text-xs">
        <svg class="w-2 h-2 fill-current opacity-30" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 2048 2048"><path d="M0 512h2048L1024 1536 0 512z"/></svg>
        <span class="opacity-70">Min: ${coin.low_24h.toLocaleString()}</span>
      </div>
    </div>
  </td>
  <td class="flex-shrink-0 rounded-r-lg">
    <img class="h-16 w-52 rounded-box" src={coin.chart} alt="{coin.name} price"/>
  </td>
</tr>