<script lang="ts">
    import { onMount } from 'svelte';
    import { showAboutDialog } from '../../lib/store';
    import { getAppInfo } from '../../lib/bridge';
    import Dialog from '../../lib/components/Dialog.svelte';
    import type { App } from '../../lib/types';

    let appInfo: App | null = $state(null);
    let loading = $state(true);
    let result: any = $state(null);

    function closeDialog() {
        showAboutDialog.set(false);
    }

    onMount(async () => {
        result = await getAppInfo();
        if (result.success && result.data) {
            appInfo = result.data;
        }
        loading = false;
    });
</script>

{#if $showAboutDialog}
    <Dialog 
        title="About"
        onClose={closeDialog}
        onConfirm={closeDialog}
    >
        {#snippet footer()}
            <div class="flex justify-end px-6 py-4 border-t border-gray-200">
                <button
                    onclick={closeDialog}
                    class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
                >
                    Confirm
                </button>
            </div>
        {/snippet}
        
        <div class="space-y-4 max-h-[60vh] overflow-y-auto">
            {#if loading}
                <div class="text-center py-4">
                    <svg class="animate-spin h-8 w-8 text-blue-600 mx-auto" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                    <p class="mt-2 text-gray-600">Loading...</p>
                </div>
            {:else if appInfo}
                <div class="text-center">
                    <div class="mx-auto bg-gray-200 border-2 border-dashed rounded-xl w-16 h-16 flex items-center justify-center mb-4">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
                        </svg>
                    </div>
                    <h3 class="text-2xl font-bold text-gray-800">{appInfo.name}</h3>
                    <p class="text-gray-600">Version {appInfo.version}</p>
                </div>

                <div class="space-y-3">
                    <div>
                        <h4 class="font-medium text-gray-700">Author</h4>
                        <p class="text-gray-600">{appInfo.author}</p>
                    </div>

                    <div>
                        <h4 class="font-medium text-gray-700">Source Code</h4>
                        <a href={appInfo.source} target="_blank" class="text-blue-600 hover:underline">{appInfo.source}</a>
                    </div>

                    <div>
                        <h4 class="font-medium text-gray-700">License</h4>
                        <a href={appInfo.license.url} target="_blank" class="text-blue-600 hover:underline">{appInfo.license.name}</a>
                    </div>

                    <div>
                        <h4 class="font-medium text-gray-700">Description</h4>
                        <p class="text-gray-600">{appInfo.description}</p>
                    </div>
                </div>
            {:else}
                <div class="text-center py-8">
                    <div class="mx-auto bg-red-100 rounded-full p-3 w-16 h-16 flex items-center justify-center mb-4">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-red-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                        </svg>
                    </div>
                    <h3 class="text-xl font-bold text-gray-800 mb-2">Failed to load app information</h3>
                    <p class="text-gray-600">{result?.msg || 'An unknown error occurred while loading app information.'}</p>
                </div>
            {/if}
        </div>
    </Dialog>
{/if}