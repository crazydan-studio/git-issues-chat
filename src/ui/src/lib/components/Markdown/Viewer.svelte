<script lang="ts">
  import { onMount } from 'svelte';
  import createCarta from './config';
  
  let { value = '' }: { value?: string } = $props();
  let html = $state('');
  let carta = $state(createCarta());
  
  $effect(() => {
    async function renderMarkdown() {
      if (value) {
        html = await carta.render(value);
      } else {
        html = '';
      }
    }
    
    renderMarkdown();
  });
</script>

<div>{@html html}</div>