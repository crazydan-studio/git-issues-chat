<script lang="ts">
    import { showAddPlatformDialog } from '../../lib/store';
    import { saveGitPlatform } from '../../lib/bridge';
    import { showNotification } from '../../lib/store';
    import Dialog from '../../lib/components/Dialog.svelte';

    let name = $state('');
    let url = $state('');
    let icon = $state('');
    let description = $state('');
    let isSaving = $state(false);

    function closeDialog() {
        showAddPlatformDialog.set(false);
        // Reset form
        name = '';
        url = '';
        icon = '';
        description = '';
    }

    async function savePlatform() {
        if (!name || !url) {
            showNotification('error', 'Platform name and URL are required');
            return;
        }

        isSaving = true;
        const result = await saveGitPlatform({
            name,
            url,
            icon,
            description
        });

        if (result.success) {
            showNotification('success', 'Platform added successfully');
            closeDialog();
        } else {
            showNotification('error', result.msg || 'Failed to add platform');
        }
        isSaving = false;
    }
</script>

{#if $showAddPlatformDialog}
    <Dialog 
        title="Add Git Platform"
        onClose={closeDialog}
        onConfirm={savePlatform}
    >
        <div class="space-y-4">
            <div>
                <label for="platformName" class="block text-sm font-medium text-gray-700 mb-1">Platform Name *</label>
                <input
                    id="platformName"
                    type="text"
                    bind:value={name}
                    class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                    placeholder="Enter platform name"
                />
            </div>

            <div>
                <label for="platformUrl" class="block text-sm font-medium text-gray-700 mb-1">URL Address *</label>
                <input
                    id="platformUrl"
                    type="text"
                    bind:value={url}
                    class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                    placeholder="Enter platform URL"
                />
            </div>

            <div>
                <label for="platformIcon" class="block text-sm font-medium text-gray-700 mb-1">Icon</label>
                <input
                    id="platformIcon"
                    type="text"
                    bind:value={icon}
                    class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                    placeholder="Enter icon identifier"
                />
            </div>

            <div>
                <label for="platformDescription" class="block text-sm font-medium text-gray-700 mb-1">Description (Optional)</label>
                <textarea
                    id="platformDescription"
                    bind:value={description}
                    class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                    placeholder="Enter platform description"
                    rows="3"
                ></textarea>
            </div>
        </div>
    </Dialog>
{/if}