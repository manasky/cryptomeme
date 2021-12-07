<script>
  import '../global.css'
  import Navbar from '@components/Navbar.svelte'
  import Footer from '@components/Footer.svelte'
  import Scripts from '@components/Scripts.svelte'
  import PageTransitions from '@components/PageTransitions.svelte';
  export let key
</script>
<script context="module">
  export const load = async ({ page }) => ({
    props: {
      key: page.path,
    },
  })
</script>
<Scripts/>
<svelte:head>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Cousine:wght@400;700&family=Nunito+Sans:wght@400;700;900&display=swap" rel="stylesheet">
</svelte:head>
<main class="w-full max-w-5xl mx-auto">
  <Navbar/>
  <div class="mt-4"></div>
  <div class="px-3">
      <PageTransitions refresh={key}>
        <slot />
      </PageTransitions>
    </div>
    <div class="mt-4"></div>
    <Footer/>
  </main>
<div class="fixed z-20 inline-block bottom-3 right-3"><a href="https://www.producthunt.com/posts/cryptomeme-wtf?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-cryptomeme-wtf" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=322126&amp;theme=dark" alt="CryptoMeme.WTF - Crypto live price + generated Memes! | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54"></a></div>
