<script lang="ts">
import type { ClassValue } from 'svelte/elements';
import type { AppUser } from '../../lib/types';
import { gitIssueCommentList, selectedGitIssue, appUser } from '../../lib/store';
import { saveGitIssueComment } from '../../lib/bridge';
import { showNotification } from '../../lib/store';
import { formatEpocMillis } from '../../lib/utils';

let { class: className }: { class?: ClassValue } = $props();

let newComment = $state('');
let isSending = $state(false);
let currentUser = $state<AppUser | null>(null);

$effect(() => {
    currentUser = $appUser;
});

async function sendComment() {
    if (!newComment.trim() || !$selectedGitIssue) return;

    isSending = true;
    const result = await saveGitIssueComment({
        issueId: $selectedGitIssue.id,
        content: newComment
    });

    if (result.success) {
        // In a real app, we would add the comment to the list
        // For now, we'll just show a success message
        showNotification('success', 'Comment sent successfully');
        newComment = '';
    } else {
        showNotification('error', result.msg || 'Failed to send comment');
    }
    isSending = false;
}

function handleKeyPress(event: KeyboardEvent) {
    if (event.key === 'Enter' && (event.ctrlKey || event.metaKey)) {
        sendComment();
    }
}
</script>

<div class={['flex flex-col h-full', className]}>
    <div class="flex-1 overflow-y-auto p-4 space-y-4">
        {#if $gitIssueCommentList.length === 0}
            <div class="text-center text-gray-500 py-8">
                No comments yet
            </div>
        {:else}
            {#each $gitIssueCommentList as comment (comment.id)}
                {#if comment.createdBy.username === currentUser?.username}
                    <!-- Current user's comment - right aligned -->
                    <div class="flex justify-end">
                        <div class="bg-blue-500 text-white rounded-lg p-3 shadow-sm max-w-[80%]">
                            <div class="flex items-center mb-2 justify-end">
                                <div>
                                    <p class="text-sm font-medium">{comment.createdBy.displayName} (You)</p>
                                    <p class="text-xs opacity-80">{formatEpocMillis(comment.createdAt, 'yyyy-MM-dd HH:mm:ss')}</p>
                                </div>
                                <div class="bg-gray-200 border-2 border-dashed rounded-xl w-8 h-8 flex items-center justify-center ml-2">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                                    </svg>
                                </div>
                            </div>
                            <p>{comment.content}</p>
                        </div>
                    </div>
                {:else}
                    <!-- Other participants' comments - left aligned -->
                    <div class="flex justify-start">
                        <div class="bg-white rounded-lg p-3 shadow-sm max-w-[80%]">
                            <div class="flex items-center mb-2">
                                <div class="bg-gray-200 border-2 border-dashed rounded-xl w-8 h-8 flex items-center justify-center mr-2">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                                    </svg>
                                </div>
                                <div>
                                    <p class="text-sm font-medium text-gray-900">{comment.createdBy.displayName}</p>
                                    <p class="text-xs text-gray-500">{formatEpocMillis(comment.createdAt, 'yyyy-MM-dd HH:mm:ss')}</p>
                                </div>
                            </div>
                            <p class="text-gray-700">{comment.content}</p>
                        </div>
                    </div>
                {/if}
            {/each}
        {/if}
    </div>

    <div class="border-t border-gray-200 p-4">
        <div class="relative">
            <textarea
                bind:value={newComment}
                onkeydown={handleKeyPress}
                placeholder="Type a comment... (Ctrl+Enter to send)"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 resize-none"
                rows="3"
            ></textarea>
            <div class="absolute bottom-2 right-2">
                <button
                    onclick={sendComment}
                    disabled={!newComment.trim() || isSending}
                    class="px-3 py-1 bg-blue-600 text-white text-sm rounded-md hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed flex items-center"
                >
                    {#if isSending}
                        <svg class="animate-spin h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                            <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                            <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                        </svg>
                    {:else}
                        Send
                    {/if}
                </button>
            </div>
        </div>
    </div>
</div>