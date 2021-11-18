import { writable } from 'svelte/store';

export let navLinks = writable([
  {
    text: 'Home',
    href: '/'
  },
]);
