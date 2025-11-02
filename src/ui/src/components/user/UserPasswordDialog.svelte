<script lang="ts">
    import { appUser, showUserPasswordDialog } from '../../lib/store';
    import { updateAppUserPassword } from '../../lib/bridge';
    import { showNotification } from '../../lib/store';
    import Dialog from '../../lib/components/Dialog.svelte';

    let oldPassword = $state('');
    let newPassword = $state('');
    let newPasswordConfirm = $state('');
    let isUpdating = $state(false);

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
        isUpdating = false;
    }
</script>

{#if $showUserPasswordDialog}
    <Dialog 
        title="Change Password"
        onClose={closeDialog}
        onConfirm={updatePassword}
    >
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
    </Dialog>
{/if}