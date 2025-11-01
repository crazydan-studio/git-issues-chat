<script lang="ts">
    import type { ClassValue } from 'svelte/elements';
    import type { Snippet } from 'svelte';
    import { onMount, onDestroy } from 'svelte';

    // Props
    let { 
        title = '',
        disableEscape = false,
        onClose,
        onConfirm,
        header,
        footer,
        children,
        class: className
    }: { 
        title?: string; 
        disableEscape?: boolean;
        onClose?: () => void;
        onConfirm?: () => void;
        header?: Snippet;
        footer?: Snippet;
        children?: Snippet;
        class?: ClassValue;
    } = $props();

    // Handle escape key
    function handleKeyDown(event: KeyboardEvent) {
        if (!disableEscape && event.key === 'Escape') {
            onClose?.();
        }
    }

    // Add event listener when component mounts
    onMount(() => {
        document.addEventListener('keydown', handleKeyDown);
    });

    // Clean up event listener when component destroys
    onDestroy(() => {
        document.removeEventListener('keydown', handleKeyDown);
    });
</script>

<div class={['fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center p-4 z-50', className]}>
    <div class="bg-white rounded-lg shadow-xl w-full max-w-md">
        <!-- Header -->
        {#if header}
            {@render header()}
        {:else}
            <div class="flex justify-between items-center px-6 py-4 border-b border-gray-200">
                <h2 class="text-xl font-bold text-gray-800">{title}</h2>
                {#if onClose}
                    <button 
                        onclick={onClose}
                        class="text-gray-500 hover:text-gray-700"
                        aria-label="Close"
                    >
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                        </svg>
                    </button>
                {/if}
            </div>
        {/if}
        
        <!-- Body -->
        <div class="p-6">
            {@render children()}
        </div>
        
        <!-- Footer -->
        {#if footer}
            {@render footer()}
        {:else if onClose || onConfirm}
            <div class="flex justify-end space-x-3 px-6 py-4 border-t border-gray-200">
                {#if onClose}
                    <button
                        onclick={onClose}
                        class="px-4 py-2 text-gray-700 bg-gray-100 rounded-md hover:bg-gray-200 transition-colors"
                    >
                        Cancel
                    </button>
                {/if}
                {#if onConfirm}
                    <button
                        onclick={onConfirm}
                        class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
                    >
                        Confirm
                    </button>
                {/if}
            </div>
        {/if}
    </div>
</div>