<script lang="ts">
    import { appUser, showUserProfileDialog } from '../lib/store.js';
    import { saveAppUserInfo } from '../lib/bridge.js';
    import { showNotification } from '../lib/store.js';

    let displayName = '';
    let avatar = '';
    let isSaving = false;

    $: if ($appUser) {
        displayName = $appUser.displayName || '';
        avatar = $appUser.avatar || '';
    }

    function closeDialog() {
        showUserProfileDialog.set(false);
    }

    async function saveProfile() {
        if (!$appUser) return;

        isSaving = true;
        try {
            const result = await saveAppUserInfo({
                id: $appUser.id,
                displayName,
                avatar
            });

            if (result.success) {
                showNotification('success', 'Profile updated successfully');
                closeDialog();
            } else {
                showNotification('error', result.msg || 'Failed to update profile');
            }
        } catch (error: any) {
            showNotification('error', 'Failed to update profile: ' + error.message);
        } finally {
            isSaving = false;
        }
    }
</script>

{#if $showUserProfileDialog}
    <div class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center p-4 z-50">
        <div class="bg-white rounded-lg shadow-xl w-full max-w-md">
            <div class="p-6">
                <div class="flex justify-between items-center mb-4">
                    <h2 class="text-xl font-bold text-gray-800">Edit Profile</h2>
                    <button 
                        on:click={closeDialog}
                        class="text-gray-500 hover:text-gray-700"
                        aria-label="Close"
                    >
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                        </svg>
                    </button>
                </div>

                <div class="space-y-4">
                    <div class="flex flex-col items-center">
                        <div class="bg-gray-200 border-2 border-dashed rounded-xl w-16 h-16 flex items-center justify-center mb-2">
                            <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                            </svg>
                        </div>
                        <p class="text-sm text-gray-600">Avatar</p>
                    </div>

                    <div>
                        <label for="displayName" class="block text-sm font-medium text-gray-700 mb-1">Display Name</label>
                        <input
                            id="displayName"
                            type="text"
                            bind:value={displayName}
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="Enter your display name"
                        />
                    </div>
                </div>

                <div class="mt-6 flex justify-end space-x-3">
                    <button
                        on:click={closeDialog}
                        class="px-4 py-2 text-gray-700 bg-gray-100 rounded-md hover:bg-gray-200 transition-colors"
                    >
                        Cancel
                    </button>
                    <button
                        on:click={saveProfile}
                        disabled={isSaving}
                        class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center"
                    >
                        {#if isSaving}
                            <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                            </svg>
                            Saving...
                        {:else}
                            Save
                        {/if}
                    </button>
                </div>
            </div>
        </div>
    </div>
{/if}