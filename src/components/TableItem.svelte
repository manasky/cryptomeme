<script>
  export let coin
</script>

<tr class="hover">
  <th class="rounded-l-lg flex-shrink-0 bg-primary">
    <div class="flex gap-3 items-center">
      <img class="w-8 h-8 rounded" src={coin.image} alt={coin.name}/>
      <div class="flex flex-col pr-4">
        <div class="font-semibold">{coin.name}</div>
        <div class="font-bold text-base-content text-opacity-50 text-xs font-mono uppercase">{coin.symbol}</div>
      </div>
    </div>
  </th>
  <td>
    <div class="flex gap-1 items-center">
      <svg class="text-success fill-current w-5 h-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M17.92,6.62a1,1,0,0,0-.54-.54A1,1,0,0,0,17,6H7A1,1,0,0,0,7,8h7.59l-8.3,8.29a1,1,0,0,0,0,1.42,1,1,0,0,0,1.42,0L16,9.41V17a1,1,0,0,0,2,0V7A1,1,0,0,0,17.92,6.62Z"/></svg>
      <span class="font-mono text-lg">${coin.current_price.toLocaleString()}</span>
    </div>
  </td>
  <td>
    <div class="flex gap-1 items-center">
      <svg class="text-success fill-current w-4 h-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M17.92,6.62a1,1,0,0,0-.54-.54A1,1,0,0,0,17,6H7A1,1,0,0,0,7,8h7.59l-8.3,8.29a1,1,0,0,0,0,1.42,1,1,0,0,0,1.42,0L16,9.41V17a1,1,0,0,0,2,0V7A1,1,0,0,0,17.92,6.62Z"/></svg>
      <span class="font-mono text-sm">{Number(Math.abs(coin.price_change_percentage_24h)).toFixed(2)}%</span>
    </div>
  </td>
  <td>
    <div class="flex gap-1 items-center">
      <svg class="text-error fill-current w-4 h-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24"><path d="M17,6a1,1,0,0,0-1,1v7.59L7.71,6.29A1,1,0,0,0,6.29,7.71L14.59,16H7a1,1,0,0,0,0,2H17a1,1,0,0,0,.38-.08,1,1,0,0,0,.54-.54A1,1,0,0,0,18,17V7A1,1,0,0,0,17,6Z"/></svg>
      <span class="font-mono text-sm">??%</span>
    </div>
  </td>
  <td class="flex-shrink-0 rounded-r-lg">
    <img src="https://s3.coinmarketcap.com/generated/sparklines/web/7d/2781/1.svg" alt="Bitcoin"/>
  </td>
</tr>