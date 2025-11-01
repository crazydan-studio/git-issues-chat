<script lang="ts">
import type { ClassValue } from 'svelte/elements';
import { gitPlatformList, selectedGitPlatform, updateSelectedGitPlatform } from '../../lib/store.js';

let { class: className, onAddPlatform }: { class?: ClassValue; onAddPlatform: () => void } = $props();

function selectPlatform(platform) {
    updateSelectedGitPlatform(platform);
}

function addPlatform() {
    onAddPlatform();
}
</script>

<div class={['flex flex-col h-full', className]}>
    <div class="px-4 py-3 border-b border-gray-200 flex items-center justify-between">
        <h2 class="font-semibold text-gray-800">Platforms</h2>
        <button 
            onclick={addPlatform}
            class="p-1 text-gray-500 hover:text-gray-700 rounded-full hover:bg-gray-100"
            aria-label="Add platform"
        >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z" clip-rule="evenodd" />
            </svg>
        </button>
    </div>
    
    <div class="flex-1 overflow-y-auto">
        {#if $gitPlatformList.length === 0}
            <div class="p-4 text-center text-gray-500">
                No platforms added yet
            </div>
        {:else}
            <ul class="divide-y divide-gray-200">
                {#each $gitPlatformList as platform (platform.id)}
                    <li>
                        <button
                            onclick={() => selectPlatform(platform)}
                            class="w-full px-4 py-3 text-left hover:bg-gray-50 flex items-center space-x-3 {($selectedGitPlatform?.id === platform.id) ? 'bg-blue-50 border-l-4 border-blue-500' : ''}"
                        >
                            <div class="bg-gray-200 border-2 border-dashed rounded-xl w-8 h-8 flex items-center justify-center">
                                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9" />
                                </svg>
                            </div>
                            <div class="flex-1 min-w-0">
                                <p class="text-sm font-medium text-gray-900 truncate">{platform.name}</p>
                                <p class="text-xs text-gray-500 truncate">{platform.url}</p>
                            </div>
                        </button>
                    </li>
                {/each}
            </ul>
        {/if}
    </div>
</div>