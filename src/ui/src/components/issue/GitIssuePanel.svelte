<script lang="ts">
import type { ClassValue } from 'svelte/elements';
import { gitIssueList, selectedGitRepo, selectedGitIssue, gitIssueCommentList, gitIssueParticipantList, showAddIssueDialog } from '../../lib/store.js';
import { getGitIssueList, getGitIssueCommentList, getGitIssueParticipantList } from '../../lib/bridge.js';
import GitIssueList from './GitIssueList.svelte';
import GitIssueCommentPanel from './GitIssueCommentPanel.svelte';
import GitIssueAddDialog from './GitIssueAddDialog.svelte';

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
    const result = await getGitIssueList({ repoId });
    if (result.success) {
        gitIssueList.set(result.data);
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
    <GitIssueList onAddIssue={openAddIssue} class="w-1/5" />
    <div class="flex-1 flex">
        <GitIssueCommentPanel />
    </div>
    
    <GitIssueAddDialog />
</div>