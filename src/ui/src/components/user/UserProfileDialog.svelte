<script lang="ts">
    import { appUser, showUserProfileDialog } from '../../lib/store';
    import { saveAppUserInfo } from '../../lib/bridge';
    import { showNotification } from '../../lib/store';
    import Dialog from '../../lib/components/Dialog.svelte';

    let displayName = $state('');
    let avatar = $state('');
    let isSaving = $state(false);

    $effect(() => {
        if ($appUser) {
            displayName = $appUser.displayName || '';
            avatar = $appUser.avatar || '';
        }
    });

    function closeDialog() {
        showUserProfileDialog.set(false);
    }

    async function saveProfile() {
        if (!$appUser) return;

        isSaving = true;
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
        isSaving = false;
    }
</script>

{#if $showUserProfileDialog}
    <Dialog 
        title="Edit Profile"
        onClose={closeDialog}
        onConfirm={saveProfile}
    >
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
    </Dialog>
{/if}