<script lang="ts">
    import type { ClassValue } from 'svelte/elements';
    import { appUserActionLogList, appUser } from '../../lib/store';
    import { getAppUserActionLogList } from '../../lib/bridge';
    import { formatEpocMillis } from '../../lib/utils';

    let { class: className, onBack }: { class?: ClassValue; onBack: () => void } = $props();
    let loading = $state(true);
    let error: string | null = $state(null);

    $effect(() => {
        (async () => {
            if ($appUser) {
                loading = true;
                error = null;
                const result = await getAppUserActionLogList({ userId: $appUser.id });
                if (result.success && result.data) {
                    appUserActionLogList.set(result.data);
                } else {
                    error = result.msg || 'Failed to load action logs';
                }
                loading = false;
            }
        })();
    });
</script>

<div class={['flex flex-col h-full overflow-hidden flex-1', className]}>
    <div class="px-4 py-3 border-b border-gray-200 flex items-center">
        <button 
            onclick={onBack}
            class="p-2 rounded-md hover:bg-gray-100 mr-2"
            aria-label="Back to chat"
        >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
            </svg>
        </button>
        <h2 class="text-lg font-semibold text-gray-800">Action Log</h2>
    </div>
    
    <div class="flex-1 overflow-y-auto p-4 flex justify-center">
        {#if loading}
            <div class="text-center py-8">
                <svg class="animate-spin h-8 w-8 text-blue-600 mx-auto" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                    <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                <p class="mt-2 text-gray-600">Loading action logs...</p>
            </div>
        {:else if error}
            <div class="text-center py-8">
                <div class="mx-auto bg-red-100 rounded-full p-3 w-16 h-16 flex items-center justify-center mb-4">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-red-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                </div>
                <h3 class="text-xl font-bold text-gray-800 mb-2">Failed to load action logs</h3>
                <p class="text-gray-600">{error}</p>
            </div>
        {:else if $appUserActionLogList.length === 0}
            <div class="text-center text-gray-500 py-8">
                No action logs found
            </div>
        {:else}
            <div class="space-y-4 w-1/2">
                {#each $appUserActionLogList as log (log.id)}
                    <div class="bg-white rounded-lg p-4 shadow-md">
                        <div class="flex justify-between items-start">
                            <div class="flex-1">
                                <p class="text-gray-800">{log.content}</p>
                                <p class="text-xs text-gray-500 mt-1">{formatEpocMillis(log.createdAt, 'yyyy-MM-dd HH:mm:ss')}</p>
                            </div>
                            <span class={`px-2 py-1 rounded-full text-xs font-medium ${
                                log.status === 'success' 
                                    ? 'bg-green-100 text-green-800' 
                                    : 'bg-red-100 text-red-800'
                            }`}>
                                {log.status === 'success' ? 'Success' : 'Error'}
                            </span>
                        </div>
                    </div>
                {/each}
            </div>
        {/if}
    </div>
</div>