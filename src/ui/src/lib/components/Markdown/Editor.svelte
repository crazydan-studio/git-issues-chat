<script lang="ts">
  import { MarkdownEditor } from 'carta-md';
  import createCarta from './config';
  import './github.css';
  
  let { value = $bindable(''), insertText: insertTextProp = '' }: { value?: string, insertText?: string } = $props();
  let carta = $state(createCarta());
  
  $effect(() => {
    if (insertTextProp && carta.input) {
      // Get current selection
      const selection = carta.input.getSelection();
      // Insert text at the current cursor position or replace selection
      carta.input.insertAt(selection.start, insertTextProp);
    }
  });
</script>

<MarkdownEditor bind:value mode="tabs" theme="github" {carta} />