<script lang="ts">
import { onMount } from 'svelte';
import type { ClassValue } from 'svelte/elements';
import { gitPlatformList, selectedGitPlatform, gitRepoList, selectedGitRepo, showAddPlatformDialog, showAddRepoDialog } from '../../lib/store';
import { getGitPlatformList, getGitRepoList } from '../../lib/bridge';
import GitPlatformList from './GitPlatformList.svelte';
import GitRepoList from './GitRepoList.svelte';
import GitPlatformAddDialog from './GitPlatformAddDialog.svelte';
import GitRepoAddDialog from './GitRepoAddDialog.svelte';

let { class: className }: { class?: ClassValue } = $props();

// Load platform list on mount
onMount(async () => {
    const result = await getGitPlatformList();
    if (result.success && result.data) {
        gitPlatformList.set(result.data);
    }
});

// Watch for selected platform changes
$effect(() => {
    if ($selectedGitPlatform) {
        loadRepoList($selectedGitPlatform.id);
    }
});

async function loadRepoList(platformId: string) {
    const result = await getGitRepoList({ platformId });
    if (result.success && result.data) {
        gitRepoList.set(result.data);
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
        <GitPlatformList onAddPlatform={openAddPlatform} />
    </div>
    <div class="flex-1 overflow-hidden">
        <GitRepoList onAddRepo={openAddRepo} />
    </div>
    
    <GitPlatformAddDialog />
    <GitRepoAddDialog />
</div>