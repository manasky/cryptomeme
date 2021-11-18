<script>
import { onMount } from 'svelte'
import { themeChange } from 'theme-change'
import ThemeItem from '@components/ThemeItem.svelte'
onMount(() => {
  themeChange(false)
})
</script>
<nav class="navbar">
  <div class="navbar-start">
    <a href="/" class="text-xl font-black link link-hover">
      CoinPrice.WTF
    </a>
  </div>
  <div class="navbar-center">
  </div>
  <div class="navbar-end gap-4">
    <a class="link link-hover" href="/">
      About
    </a>
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