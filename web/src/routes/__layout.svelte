<script>
  import '../global.css'
  import Navbar from '@components/Navbar.svelte'
  import Footer from '@components/Footer.svelte'
  import Scripts from '@components/Scripts.svelte'
</script>
<Scripts/>
<svelte:head>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Cousine:wght@400;700&family=Nunito+Sans:wght@400;700;900&display=swap" rel="stylesheet">
</svelte:head>
<main class="w-full max-w-3xl mx-auto">
  <Navbar/>
  <div class="mt-4"></div>
  <div class="px-3">
    <slot />
  </div>
  <div class="mt-4"></div>
  <Footer/>
</main>