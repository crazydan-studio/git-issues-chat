<script lang="ts">
import type { ClassValue } from 'svelte/elements';
import { gitPlatformRepoList, selectedGitPlatform, selectedGitRepo, updateSelectedGitRepo } from '../lib/store.js';
import { createEventDispatcher } from 'svelte';

let { class: className }: { class?: ClassValue } = $props();

const dispatch = createEventDispatcher();

function selectRepo(repo) {
    updateSelectedGitRepo(repo);
}

function addRepo() {
    dispatch('addRepo');
}
</script>

<div class={['flex flex-col h-full border-t border-gray-200', className]}>
    <div class="px-4 py-3 border-b border-gray-200 flex items-center justify-between">
        <h2 class="font-semibold text-gray-800">Repositories</h2>
        <button 
            on:click={addRepo}
            disabled={!$selectedGitPlatform}
            class="p-1 text-gray-500 hover:text-gray-700 rounded-full hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
            aria-label="Add repository"
        >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z" clip-rule="evenodd" />
            </svg>
        </button>
    </div>
    
    <div class="flex-1 overflow-y-auto">
        {#if !$selectedGitPlatform}
            <div class="p-4 text-center text-gray-500">
                Select a platform first
            </div>
        {:else if $gitPlatformRepoList.length === 0}
            <div class="p-4 text-center text-gray-500">
                No repositories added yet
            </div>
        {:else}
            <ul class="divide-y divide-gray-200">
                {#each $gitPlatformRepoList as repo (repo.id)}
                    <li>
                        <button
                            on:click={() => selectRepo(repo)}
                            class="w-full px-4 py-3 text-left hover:bg-gray-50 flex items-center space-x-3 {($selectedGitRepo?.id === repo.id) ? 'bg-blue-50 border-l-4 border-blue-500' : ''}"
                        >
                            <div class="bg-gray-200 border-2 border-dashed rounded-xl w-8 h-8 flex items-center justify-center">
                                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15a4 4 0 004 4h9a5 5 0 10-.1-9.999 5.002 5.002 0 10-9.78 2.096A4 4 0 003 15z" />
                                </svg>
                            </div>
                            <div class="flex-1 min-w-0">
                                <p class="text-sm font-medium text-gray-900 truncate">{repo.name}</p>
                                <p class="text-xs text-gray-500 truncate">{repo.description || 'No description'}</p>
                            </div>
                        </button>
                    </li>
                {/each}
            </ul>
        {/if}
    </div>
</div>