<script lang="ts">
import { onMount } from 'svelte';
import type { ClassValue } from 'svelte/elements';
import { gitPlatformList, selectedGitPlatform, gitPlatformRepoList, selectedGitRepo, showAddPlatformDialog, showAddRepoDialog } from '../../lib/store.js';
import { getGitPlatformList, getGitPlatformRepoList } from '../../lib/bridge.js';
import GitPlatformList from './GitPlatformList.svelte';
import GitRepoList from './GitRepoList.svelte';

let { class: className }: { class?: ClassValue } = $props();

// Load platform list on mount
onMount(async () => {
    const result = await getGitPlatformList();
    if (result.success) {
        gitPlatformList.set(result.data);
    }
});

// Watch for selected platform changes
$effect(() => {
    if ($selectedGitPlatform) {
        loadRepoList($selectedGitPlatform.id);
    }
});

async function loadRepoList(platformId) {
    const result = await getGitPlatformRepoList({ platformId });
    if (result.success) {
        gitPlatformRepoList.set(result.data);
    }
}

function openAddPlatform() {
    showAddPlatformDialog.set(true);
}

function openAddRepo() {
    showAddRepoDialog.set(true);
}
</script>

<div class={['flex flex-col h-full border-r border-gray-200', className]}>
    <div class="h-1/3 overflow-hidden">
        <GitPlatformList on:addPlatform={openAddPlatform} />
    </div>
    <div class="flex-1 overflow-hidden">
        <GitRepoList on:addRepo={openAddRepo} />
    </div>
</div>