<script lang="ts">
import type { ClassValue } from 'svelte/elements';
import { gitRepoIssueList, selectedGitRepo, selectedGitIssue, updateSelectedGitIssue } from '../lib/store.js';
import { createEventDispatcher } from 'svelte';

let { class: className }: { class?: ClassValue } = $props();

const dispatch = createEventDispatcher();

function selectIssue(issue) {
    updateSelectedGitIssue(issue);
}

function addIssue() {
    dispatch('addIssue');
}
</script>

<div class={['flex flex-col border-r border-gray-200', className]}>
    <div class="px-4 py-3 border-b border-gray-200 flex items-center justify-between">
        <h2 class="font-semibold text-gray-800">Issues</h2>
        <button 
            on:click={addIssue}
            disabled={!$selectedGitRepo}
            class="p-1 text-gray-500 hover:text-gray-700 rounded-full hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed"
            aria-label="Add issue"
        >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z" clip-rule="evenodd" />
            </svg>
        </button>
    </div>
    
    <div class="flex-1 overflow-y-auto">
        {#if !$selectedGitRepo}
            <div class="p-4 text-center text-gray-500">
                Select a repository first
            </div>
        {:else if $gitRepoIssueList.length === 0}
            <div class="p-4 text-center text-gray-500">
                No issues found
            </div>
        {:else}
            <ul class="divide-y divide-gray-200">
                {#each $gitRepoIssueList as issue (issue.id)}
                    <li>
                        <button
                            on:click={() => selectIssue(issue)}
                            class="w-full px-4 py-3 text-left hover:bg-gray-50 {($selectedGitIssue?.id === issue.id) ? 'bg-blue-50 border-l-4 border-blue-500' : ''}"
                        >
                            <div class="flex justify-between">
                                <p class="text-sm font-medium text-gray-900 truncate">#{issue.code} {issue.title}</p>
                                <span class="text-xs text-gray-500">{issue.createdAt.split(' ')[0]}</span>
                            </div>
                            <div class="flex items-center mt-1">
                                <div class="bg-gray-200 border-2 border-dashed rounded-xl w-5 h-5 flex items-center justify-center mr-2">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                                    </svg>
                                </div>
                                <p class="text-xs text-gray-500 truncate">{issue.author.displayName}</p>
                            </div>
                        </button>
                    </li>
                {/each}
            </ul>
        {/if}
    </div>
</div>