<script lang="ts">
    import { appUser, showUserPasswordDialog } from '../lib/store.js';
    import { updateAppUserPassword } from '../lib/bridge.js';
    import { showNotification } from '../lib/store.js';

    let oldPassword = '';
    let newPassword = '';
    let newPasswordConfirm = '';
    let isUpdating = false;

    function closeDialog() {
        showUserPasswordDialog.set(false);
        // Reset form
        oldPassword = '';
        newPassword = '';
        newPasswordConfirm = '';
    }

    async function updatePassword() {
        if (!oldPassword || !newPassword || !newPasswordConfirm) {
            showNotification('error', 'Please fill in all fields');
            return;
        }

        if (newPassword !== newPasswordConfirm) {
            showNotification('error', 'New passwords do not match');
            return;
        }

        if (newPassword.length < 6) {
            showNotification('error', 'Password must be at least 6 characters');
            return;
        }

        if (!$appUser) return;

        isUpdating = true;
        try {
            const result = await updateAppUserPassword({
                userId: $appUser.id,
                oldPassword,
                newPassword,
                newPasswordConfirm
            });

            if (result.success) {
                showNotification('success', 'Password updated successfully');
                closeDialog();
            } else {
                showNotification('error', result.msg || 'Failed to update password');
            }
        } catch (error: any) {
            showNotification('error', 'Failed to update password: ' + error.message);
        } finally {
            isUpdating = false;
        }
    }
</script>

{#if $showUserPasswordDialog}
    <div class="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center p-4 z-50">
        <div class="bg-white rounded-lg shadow-xl w-full max-w-md">
            <div class="p-6">
                <div class="flex justify-between items-center mb-4">
                    <h2 class="text-xl font-bold text-gray-800">Change Password</h2>
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
                    <div>
                        <label for="oldPassword" class="block text-sm font-medium text-gray-700 mb-1">Current Password</label>
                        <input
                            id="oldPassword"
                            type="password"
                            bind:value={oldPassword}
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="Enter current password"
                        />
                    </div>

                    <div>
                        <label for="newPassword" class="block text-sm font-medium text-gray-700 mb-1">New Password</label>
                        <input
                            id="newPassword"
                            type="password"
                            bind:value={newPassword}
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="Enter new password"
                        />
                    </div>

                    <div>
                        <label for="newPasswordConfirm" class="block text-sm font-medium text-gray-700 mb-1">Confirm New Password</label>
                        <input
                            id="newPasswordConfirm"
                            type="password"
                            bind:value={newPasswordConfirm}
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="Confirm new password"
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
                        on:click={updatePassword}
                        disabled={isUpdating}
                        class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center"
                    >
                        {#if isUpdating}
                            <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                            </svg>
                            Updating...
                        {:else}
                            Update
                        {/if}
                    </button>
                </div>
            </div>
        </div>
    </div>
{/if}