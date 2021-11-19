<script>
import { onMount } from 'svelte'
import { themeChange } from 'theme-change'
onMount(() => {
  themeChange(false)
})
let themes = [
  'forest',
  'dark',
  'garden',
  'cupcake',
  'light',
  'valentine',
  'synthwave',
]
</script>
<nav class="px-3 navbar">
  <div class="navbar-start">

    <div class="dropdown">
      <div tabindex="0" class="btn btn-ghost btn-square mask mask-squircle">
        <svg class="inline-block w-6 h-6 stroke-current" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path></svg>
      </div> 
      <ul tabindex="0" class="p-2 mt-1 shadow-2xl menu compact dropdown-content bg-neutral-focus text-neutral-content rounded-box w-52">
        <li>
          <a href="/">Home</a>
        </li>
        <li>
          <a href="/">About</a>
        </li>
      </ul>
    </div>

  </div>
  <div class="navbar-center">
    <a href="/" class="text-xl font-black link link-hover">
      CoinPrice.WTF
    </a>
  </div>
  <div class="gap-4 navbar-end">

    <div class="hidden gap-2 md:flex">
      {#each themes as themeName}
        <button class="w-4 h-4 border-2 rounded-full border-opacity-10 border-base-content bg-base-100" 
          data-theme={themeName} 
          data-set-theme={themeName} 
          title={themeName}
        ></button>
      {/each}
    </div>

    <div class="dropdown dropdown-end md:hidden">
      <div tabindex="0" class="btn btn-ghost btn-square mask mask-squircle">
        <svg class="inline-block w-6 h-6 stroke-current" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01"></path></svg>
      </div> 
      <div tabindex="0" class="grid grid-cols-4 gap-4 p-6 mt-1 shadow-2xl dropdown-content bg-neutral-focus text-neutral-content rounded-box w-52">
        {#each themes as themeName}
          <button class="w-6 h-6 border rounded-full border-base-content border-opacity-30 bg-base-100" 
            data-theme={themeName} 
            data-set-theme={themeName} 
            title={themeName}
          ></button>
        {/each}
      </div>
    </div>

  </div>
</nav>