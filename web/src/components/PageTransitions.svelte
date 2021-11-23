<script>
  import { fly, fade } from "svelte/transition"
  export let refresh = ""
</script>

{#key refresh}
  <div
  in:fly={{
    y: -5,
    duration: 300,
    delay: 300,
  }}
  out:fade={{
    duration: 300
  }}>
    <slot />
  </div>
{/key}