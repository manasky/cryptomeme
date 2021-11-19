<script>
import { onMount } from 'svelte'
import { themeChange } from 'theme-change'
import ThemeItem from '@components/ThemeItem.svelte'
onMount(() => {
  themeChange(false)
})
</script>
<nav class="px-3 navbar">
  <div class="navbar-start">
    <div class="dropdown">
      <div tabindex="0" class="btn btn-ghost btn-square mask mask-squircle">
        <svg class="inline-block w-6 h-6 stroke-current" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path></svg>
      </div> 
      <ul tabindex="0" class="p-2 mt-1 shadow-2xl compact menu dropdown-content bg-neutral text-neutral-content rounded-box w-52">
        <li>
          <a href="/">About</a>
        </li> 
        <li>
          <a href="/">Item 2</a>
        </li> 
        <li>
          <a href="/">Item 3</a>
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
    <div class="flex gap-2">
      <ThemeItem name="forest" />
      <ThemeItem name="dark" />
      <ThemeItem name="garden" />
      <ThemeItem name="cupcake" />
      <ThemeItem name="light" />
      <ThemeItem name="valentine" />
      <ThemeItem name="cyberpunk" />
      <ThemeItem name="synthwave" />
    </div>
  </div>
</nav>