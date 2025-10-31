<script lang="ts">
import type { ClassValue } from 'svelte/elements';
import { gitIssueParticipantList } from '../lib/store.js';

let { class: className }: { class?: ClassValue } = $props();
</script>

<div class={['flex flex-col h-full', className]}>
    <div class="px-4 py-3 border-b border-gray-200">
        <h3 class="font-semibold text-gray-800">Participants</h3>
    </div>
    
    <div class="flex-1 overflow-y-auto">
        {#if $gitIssueParticipantList.length === 0}
            <div class="p-4 text-center text-gray-500">
                No participants
            </div>
        {:else}
            <ul class="divide-y divide-gray-200">
                {#each $gitIssueParticipantList as participant (participant.username)}
                    <li>
                        <div class="block px-4 py-3 hover:bg-gray-50 flex items-center space-x-3">
                            <div class="bg-gray-200 border-2 border-dashed rounded-xl w-8 h-8 flex items-center justify-center">
                                <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                                </svg>
                            </div>
                            <div class="flex-1 min-w-0">
                                <a href={participant.url} target="_blank" class="text-sm font-medium text-blue-600 hover:text-blue-800 hover:underline truncate">{participant.displayName}</a>
                                <p class="text-xs text-gray-500 truncate">@{participant.username}</p>
                            </div>
                        </div>
                    </li>
                {/each}
            </ul>
        {/if}
    </div>
</div>