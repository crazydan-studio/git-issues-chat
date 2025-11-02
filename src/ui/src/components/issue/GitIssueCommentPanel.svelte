<script lang="ts">
    import { gitIssueCommentList, gitIssueParticipantList, selectedGitIssue } from '../../lib/store.js';
    import GitIssueCommentList from './GitIssueCommentList.svelte';
    import GitIssueParticipantList from './GitIssueParticipantList.svelte';
    import { formatEpocMillis } from '../../lib/utils.ts';
</script>

<div class="flex flex-col flex-1">
    {#if !$selectedGitIssue}
        <div class="flex-1 flex items-center justify-center text-gray-500">
            <p>Select an issue to view messages</p>
        </div>
    {:else}
        <div class="px-4 py-3 border-b border-gray-200">
            <h2 class="text-lg font-semibold text-gray-800">{$selectedGitIssue.title}</h2>
            <div class="flex items-center mt-1 text-sm text-gray-600">
                <a href={$selectedGitIssue.url} target="_blank" class="text-blue-600 hover:underline">#{ $selectedGitIssue.code }</a>
                <span class="mx-2">•</span>
                <span>{ formatEpocMillis($selectedGitIssue.createdAt, 'yyyy-MM-dd') }</span>
                <span class="mx-2">•</span>
                <div class="flex items-center">
                    <div class="bg-gray-200 border-2 border-dashed rounded-xl w-5 h-5 flex items-center justify-center mr-1">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                        </svg>
                    </div>
                    <span>{ $selectedGitIssue.createdBy.displayName }</span>
                </div>
            </div>
        </div>
        
        <div class="flex flex-1 overflow-hidden">
            <GitIssueCommentList class="w-4/5 border-r border-gray-200" />
            <GitIssueParticipantList class="w-1/5" />
        </div>
    {/if}
</div>