<script lang="ts">
import type { ClassValue } from 'svelte/elements';
import { gitRepoIssueList, selectedGitRepo, selectedGitIssue, gitIssueCommentList, gitIssueParticipantList, showAddIssueDialog } from '../lib/store.js';
import { getGitRepoIssueList, getGitIssueCommentList, getGitIssueParticipantList } from '../lib/bridge.js';
import IssueList from './IssueList.svelte';
import IssueMessagePanel from './IssueMessagePanel.svelte';

let { class: className }: { class?: ClassValue } = $props();

// Watch for selected repo changes
$effect(() => {
    if ($selectedGitRepo) {
        loadIssueList($selectedGitRepo.id);
    }
});

// Watch for selected issue changes
$effect(() => {
    if ($selectedGitIssue) {
        loadCommentList($selectedGitIssue.id);
        loadParticipantList($selectedGitIssue.id);
    }
});

async function loadIssueList(repoId) {
    const result = await getGitRepoIssueList({ repoId });
    if (result.success) {
        gitRepoIssueList.set(result.data);
    }
}

async function loadCommentList(issueId) {
    const result = await getGitIssueCommentList({ issueId });
    if (result.success) {
        gitIssueCommentList.set(result.data);
    }
}

async function loadParticipantList(issueId) {
    const result = await getGitIssueParticipantList({ issueId });
    if (result.success) {
        gitIssueParticipantList.set(result.data);
    }
}

function openAddIssue() {
    showAddIssueDialog.set(true);
}
</script>

<div class={['flex h-full', className]}>
    <IssueList on:addIssue={openAddIssue} class="w-1/5" />
    <div class="flex-1 flex">
        <IssueMessagePanel />
    </div>
</div>