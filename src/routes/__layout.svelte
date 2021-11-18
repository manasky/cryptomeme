<script>
  import '../global.css'
  import Navbar from '@components/Navbar.svelte'
  import Footer from '@components/Footer.svelte'
</script>
<svelte:head>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Cousine&family=Nunito+Sans:wght@400;700;900&display=swap" rel="stylesheet">
</svelte:head>
<main class="w-full mx-auto max-w-3xl">
  <Navbar/>
  <div class="mt-4"></div>
  <div class="px-3">
    <slot />
  </div>
  <div class="mt-10"></div>
  <Footer/>
</main>