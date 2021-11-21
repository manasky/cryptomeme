<script>
  import '../global.css'
  import Navbar from '@components/Navbar.svelte'
  import Footer from '@components/Footer.svelte'
</script>
<svelte:head>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Cousine:wght@400;700&family=Nunito+Sans:wght@400;700;900&display=swap" rel="stylesheet">
  <link rel="icon" type="image/x-icon" href="data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 16 16'><text x='0' y='14'>📈📉</text></svg>" />
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