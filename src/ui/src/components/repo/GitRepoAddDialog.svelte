<script lang="ts">
    import { showAddRepoDialog } from '../../lib/store';
    import { getGitRepoInfo, saveGitRepo } from '../../lib/bridge';
    import { Notify } from '../../lib/components/Notification';
    import Dialog from '../../lib/components/Dialog.svelte';

    // Form states: 'input' for repo input, 'confirm' for repo confirmation
    let formState = $state('input');
    let platformUrl = $state('');
    let repoName = $state('');
    let repoInfo = $state({
        name: '',
        url: '',
        icon: '',
        description: ''
    });
    let isFetching = $state(false);
    let isSaving = $state(false);

    function closeDialog() {
        showAddRepoDialog.set(false);
        // Reset form
        formState = 'input';
        platformUrl = '';
        repoName = '';
        repoInfo = {
            name: '',
            url: '',
            icon: '',
            description: ''
        };
    }

    function goBack() {
        formState = 'input';
    }

    async function fetchRepoInfo() {
        if (!platformUrl || !repoName) {
            Notify.warning('Platform URL and repository name are required');
            return;
        }

        // Validate repo name format
        if (!repoName.includes('/') || repoName.split('/').length !== 2) {
            Notify.warning('Repository name must be in the format <username>/<reponame>');
            return;
        }

        isFetching = true;
        const result = await getGitRepoInfo({
            platformId: platformUrl, // In a real app, this would be the actual platform ID
            repoName
        });

        if (result.success && result.data) {
            repoInfo = {
                name: result.data.name || repoName,
                url: result.data.url || `${platformUrl}/${repoName}`,
                icon: result.data.icon || 'repo-icon-default',
                description: result.data.description || ''
            };
            formState = 'confirm';
        } else {
            Notify.error(result.msg || 'Failed to fetch repository information');
        }
        isFetching = false;
    }

    async function saveRepo() {
        if (!repoInfo.name || !repoInfo.url) {
            Notify.warning('Repository name and URL are required');
            return;
        }

        isSaving = true;
        const result = await saveGitRepo({
            platform: platformUrl, // In a real app, this would be the actual platform ID
            name: repoName,
            url: repoInfo.url,
            icon: repoInfo.icon,
            description: repoInfo.description
        });

        if (result.success) {
            Notify.success('Repository added successfully');
            closeDialog();
        } else {
            Notify.error(result.msg || 'Failed to add repository');
        }
        isSaving = false;
    }
</script>

{#if $showAddRepoDialog}
    <Dialog 
        title="Add Git Repository"
        onClose={closeDialog}
        onConfirm={formState === 'input' ? fetchRepoInfo : saveRepo}
    >

        {#if formState === 'input'}
            <div class="space-y-4 p-6">
                <div class="flex space-x-2">
                    <div class="flex-1">
                        <label for="platformUrl" class="block text-sm font-medium text-gray-700 mb-1">Git Platform URL</label>
                        <input
                            id="platformUrl"
                            type="text"
                            bind:value={platformUrl}
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="https://github.com"
                        />
                    </div>
                    <div class="flex items-end">
                        <span class="text-gray-500 mb-2">/</span>
                    </div>
                    <div class="flex-1">
                        <label for="repoName" class="block text-sm font-medium text-gray-700 mb-1">Repository Path</label>
                        <input
                            id="repoName"
                            type="text"
                            bind:value={repoName}
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="<username>/<reponame>"
                        />
                    </div>
                </div>
            </div>
        {:else}
            <div class="p-6">

                <div class="space-y-4">
                    <div>
                        <label for="repoNameConfirm" class="block text-sm font-medium text-gray-700 mb-1">Repository Name</label>
                        <input
                            id="repoNameConfirm"
                            type="text"
                            bind:value={repoInfo.name}
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 bg-gray-100"
                            readonly
                        />
                    </div>

                    <div>
                        <label for="repoUrl" class="block text-sm font-medium text-gray-700 mb-1">Repository URL</label>
                        <input
                            id="repoUrl"
                            type="text"
                            bind:value={repoInfo.url}
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 bg-gray-100"
                            readonly
                        />
                    </div>

                    <div>
                        <label for="repoIcon" class="block text-sm font-medium text-gray-700 mb-1">Icon</label>
                        <input
                            id="repoIcon"
                            type="text"
                            bind:value={repoInfo.icon}
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                        />
                    </div>

                    <div>
                        <label for="repoDescription" class="block text-sm font-medium text-gray-700 mb-1">Description</label>
                        <textarea
                            id="repoDescription"
                            bind:value={repoInfo.description}
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="Enter repository description"
                            rows="3"
                        ></textarea>
                    </div>
                </div>
            </div>
        {/if}

        {#snippet footer()}
            <div class="flex justify-end space-x-3 px-6 py-4 border-t border-gray-200">
                <button
                    onclick={closeDialog}
                    class="px-4 py-2 text-gray-700 bg-gray-100 rounded-md hover:bg-gray-200 transition-colors"
                >
                    Cancel
                </button>
                {#if formState === 'input'}
                    <button
                        onclick={fetchRepoInfo}
                        disabled={!platformUrl || !repoName || isFetching}
                        class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center"
                    >
                        {#if isFetching}
                            <svg class="animate-spin h-4 w-4 text-white mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                            </svg>
                            Fetching...
                        {:else}
                            Next
                        {/if}
                    </button>
                {:else}
                    <button
                        onclick={goBack}
                        class="px-4 py-2 text-gray-700 bg-gray-100 rounded-md hover:bg-gray-200 transition-colors"
                    >
                        Previous
                    </button>
                    <button
                        onclick={saveRepo}
                        disabled={isSaving}
                        class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center"
                    >
                        {#if isSaving}
                            <svg class="animate-spin h-4 w-4 text-white mr-2" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                            </svg>
                            Saving...
                        {:else}
                            Confirm
                        {/if}
                    </button>
                {/if}
            </div>
        {/snippet}
    </Dialog>
{/if}

