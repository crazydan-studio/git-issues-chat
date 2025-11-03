<script lang="ts">
    import type { ClassValue } from 'svelte/elements';
    import { micromark } from 'micromark';

    // Props
    let { 
        value = $bindable(''),
        onValueChange,
        class: className
    }: { 
        value?: string; 
        onValueChange?: (value: string) => void;
        class?: ClassValue;
    } = $props();

    // State
    let mode = $state<'edit' | 'preview'>('edit');
    let textareaElement = $state<HTMLTextAreaElement | null>(null);

    // Functions
    function handleInput(event: Event) {
        const target = event.target as HTMLTextAreaElement;
        value = target.value;
        onValueChange?.(target.value);
    }

    function switchToEdit() {
        mode = 'edit';
    }

    function switchToPreview() {
        mode = 'preview';
    }

    function insertText(text: string) {
        if (textareaElement) {
            const start = textareaElement.selectionStart;
            const end = textareaElement.selectionEnd;
            const newValue = value.substring(0, start) + text + value.substring(end);
            value = newValue;
            onValueChange?.(newValue);
            
            // Set cursor position after the inserted text
            setTimeout(() => {
                if (textareaElement) {
                    textareaElement.selectionStart = start + text.length;
                    textareaElement.selectionEnd = start + text.length;
                    textareaElement.focus();
                }
            }, 0);
        }
    }

    // Export function for parent components to insert text
    export { insertText };
</script>

<div class={['border border-gray-300 rounded-lg overflow-hidden', className]}>
    <!-- Header -->
    <div class="flex border-b border-gray-300 bg-gray-50">
        <button
            type="button"
            onclick={switchToEdit}
            class={`px-4 py-2 text-sm font-medium ${
                mode === 'edit' 
                    ? 'text-blue-600 border-b-2 border-blue-600' 
                    : 'text-gray-500 hover:text-gray-700'
            }`}
        >
            Edit
        </button>
        <button
            type="button"
            onclick={switchToPreview}
            class={`px-4 py-2 text-sm font-medium ${
                mode === 'preview' 
                    ? 'text-blue-600 border-b-2 border-blue-600' 
                    : 'text-gray-500 hover:text-gray-700'
            }`}
        >
            Preview
        </button>
    </div>
    
    <!-- Body -->
    <div class="p-4 min-h-32">
        {#if mode === 'edit'}
            <textarea
                value={value}
                oninput={handleInput}
                class="w-full h-48 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                bind:this={textareaElement}
            ></textarea>
        {:else}
            <div class="prose max-w-none">
                {@html micromark(value || '')}
            </div>
        {/if}
    </div>
</div>