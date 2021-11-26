<script>
import { onMount } from 'svelte'
import Meme from '@components/Meme.svelte'

let meme = {
  "caption" : 'Is this a pump?',
  "image" : 'is-this-a-pigeon.jpg',
  "content" : [
    {
      "x": 500,
      "y": 280,
      "text": 'BTC',
      "image": '/logos/btc.png'
    }
  ]
}

function saveToLocalStorage() {
  localStorage.setItem('workshop_editor_canvas_data', JSON.stringify(meme));
}
function loadLocalStorage() {
  let workshop_editor_canvas_data = localStorage.getItem('workshop_editor_canvas_data')
  if (workshop_editor_canvas_data) {
    meme = JSON.parse(workshop_editor_canvas_data)
  }
}
function addMemeContent() {
  meme.content = [...meme.content, {
    "x": 50 + Math.floor(Math.random() * 450),
    "y": 300 + Math.floor(Math.random() * 200),
    "text": "text",
    "image": "",
  }];
}
function removeMemeContent(index) {
  meme.content.splice(index, 1)
  meme.content = meme.content;
}
onMount(() => {
  loadLocalStorage()
  let autoSave = setInterval(saveToLocalStorage,1000)
})
</script>
<svelte:head>
  <title>Editor</title>
  <meta name="robots" content="noindex">
</svelte:head>
<div class="flex flex-col-reverse gap-3 lg:flex-row">
  <div class="flex flex-col gap-3">

    <div class="form-control">
      <span class="mx-3 label-text">image</span>
      <input class="w-full input input-bordered" bind:value={meme.image}/>
    </div>
    <div class="form-control">
      <span class="mx-3 label-text">caption</span>
      <input class="w-full input input-bordered" bind:value={meme.caption}/>
    </div>
    
    {#each meme.content as name, i}
      <div class="p-4 rounded-box bg-base-200 indicator">
        
        <div>
          <div class="flex gap-2">
            <div class="form-control">
              <span class="mx-3 label-text">x</span>
              <input class="w-full input input-sm input-bordered" type="number" bind:value={meme.content[i].x}/>
            </div>
            <div class="form-control">
              <span class="mx-3 label-text">y</span>
              <input class="w-full input input-sm input-bordered" type="number" bind:value={meme.content[i].y}/>
            </div>
          </div>
          <div class="flex gap-2">
            <div class="form-control">
              <span class="mx-3 label-text">text</span>
              <input class="w-full input input-sm input-bordered" bind:value={meme.content[i].text}/>
            </div>
            <div class="form-control">
              <span class="mx-3 label-text">image</span>
              <input class="w-full input input-sm input-bordered" bind:value={meme.content[i].image}/>
            </div>
          </div>
        </div>

        <div class="indicator-item indicator-start">
          <button class="btn btn-circle btn-warning btn-xs" on:click={() => removeMemeContent(i)}>
            <svg class="w-4 h-4 fill-current" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 32 32"><path d="M 15 4 C 14.476563 4 13.941406 4.183594 13.5625 4.5625 C 13.183594 4.941406 13 5.476563 13 6 L 13 7 L 7 7 L 7 9 L 8 9 L 8 25 C 8 26.644531 9.355469 28 11 28 L 23 28 C 24.644531 28 26 26.644531 26 25 L 26 9 L 27 9 L 27 7 L 21 7 L 21 6 C 21 5.476563 20.816406 4.941406 20.4375 4.5625 C 20.058594 4.183594 19.523438 4 19 4 Z M 15 6 L 19 6 L 19 7 L 15 7 Z M 10 9 L 24 9 L 24 25 C 24 25.554688 23.554688 26 23 26 L 11 26 C 10.445313 26 10 25.554688 10 25 Z M 12 12 L 12 23 L 14 23 L 14 12 Z M 16 12 L 16 23 L 18 23 L 18 12 Z M 20 12 L 20 23 L 22 23 L 22 12 Z"/></svg>
          </button>
        </div>

      </div>
    {/each}
    <div>
      <button on:click={addMemeContent} class="btn btn-xs">add</button>
    </div>


  </div>

  <div class="max-w-[700px] flex-shrink-0 mx-auto">
    <Meme 
      coin={{ 
        name: 'Dogecoin',
        symbol: 'doge',
        current_price: '999',
        price_change_percentage_24h: '99',

        meme_caption: meme.caption,
        meme: '/memes/'+meme.image,
        meme_content: meme.content,

      }}
      classes="w-full mb-0"
      wrapperClasses="rounded-box"
    />
    <div class="w-full mt-4 overflow-x-scroll mockup-code">
      <pre class="leading-[0.2em]">
        <code class="text-xs">{`{
    "caption": "${meme.caption}",
    "image": "${'/memes/'+meme.image}",
    "content": [`}
      {#each meme.content as name, i}
      {`{
        "x": ${meme.content[i].x},
        "y": ${meme.content[i].y},
        ${meme.content[i].text? '"text": "'+meme.content[i].text+'"' : ''}${(meme.content[i].text && meme.content[i].image )? ',' : ''}
        ${meme.content[i].image? '"image": "'+meme.content[i].image+'"' : ''}
      }`}{#if i < meme.content.length - 1}{`,`}{/if}
      {/each}
    {`]
  }`}</code>
      </pre>
    </div>
  </div>
</div>