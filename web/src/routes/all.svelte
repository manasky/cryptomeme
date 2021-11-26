<script>
import Meme from '@components/Meme.svelte'
import all from '../../../dataset.json'

function replaceStrings(str,name) {
  str = str.replace('{{.Name}}', `${name}`)
  str = str.replace('${{FormatNumber .CurrentPrice}}', `$9,999`)
  return str
}

</script>
<svelte:head>
  <title>All</title>
  <meta name="robots" content="noindex">
</svelte:head>
{#each all.coins as { name, ranges }, coinIndex}
  {#each ranges as { range }, rangeIndex}
    <h2 class="flex items-center gap-1 mt-10 mb-2">
      <span class="text-lg font-black">{name}</span>
      <span class="badge badge-outline">{ranges[rangeIndex].value}</span>
    </h2>
    <div class="grid grid-cols-6 gap-2">
      {#each ranges[rangeIndex].memes as { meme }, memeIndex}
        <Meme 
          coin={{ 
            name: 'Dogecoin',
            symbol: 'doge',
            current_price: '999',
            price_change_percentage_24h: '99',

            meme_caption: replaceStrings(ranges[rangeIndex]['memes'][memeIndex].caption, name),
            meme: ranges[rangeIndex]['memes'][memeIndex].image,
            meme_content: ranges[rangeIndex]['memes'][memeIndex].content,

          }}
          classes="w-full mb-0"
          wrapperClasses="rounded-box"
        />
      {/each}
    </div>
  {/each}
  <div class="h-2 my-10 bg-base-content"></div>
{/each}