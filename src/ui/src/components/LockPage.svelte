<script lang="ts">
    import { onMount } from 'svelte';
    import { appUser } from '../lib/store.js';
    import { verifyAppUser } from '../lib/bridge.js';
    import { showNotification } from '../lib/store.js';

    let password = '';
    let isVerifying = false;
    let userId = '';

    $: if ($appUser) {
        userId = $appUser.id;
    }

    async function handleUnlock() {
        if (!password) {
            showNotification('error', 'Please enter a password');
            return;
        }

        isVerifying = true;
        try {
            const result = await verifyAppUser({ userId, password });
            if (result.success) {
                // User is already in store, no need to update
                password = ''; // Reset password field
            } else {
                showNotification('error', result.msg || 'Authentication failed');
            }
        } catch (error: any) {
            showNotification('error', 'Authentication failed: ' + error.message);
        } finally {
            isVerifying = false;
        }
    }

    function handleKeyPress(event: KeyboardEvent) {
        if (event.key === 'Enter') {
            handleUnlock();
        }
    }
</script>

<div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex flex-col">
    <div class="flex flex-1 items-center justify-center p-4">
        <div class="w-full max-w-md">
            <div class="text-center mb-8">
                <div class="mx-auto bg-gray-200 border-2 border-dashed rounded-xl w-16 h-16 flex items-center justify-center mb-4">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                    </svg>
                </div>
                <h1 class="text-3xl font-bold text-gray-800">Git Issues Chat</h1>
                <p class="text-gray-600 mt-2">Application is locked</p>
            </div>

            <div class="bg-white rounded-lg shadow-md p-6">
                <div class="space-y-4">
                    <div>
                        <label for="password" class="block text-sm font-medium text-gray-700 mb-1">Password</label>
                        <input
                            id="password"
                            type="password"
                            bind:value={password}
                            on:keydown={handleKeyPress}
                            class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                            placeholder="Enter your password to unlock"
                        />
                    </div>

                    <button
                        on:click={handleUnlock}
                        disabled={isVerifying}
                        class="w-full bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-md transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center"
                    >
                        {#if isVerifying}
                            <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                            </svg>
                            Unlocking...
                        {:else}
                            Unlock
                        {/if}
                    </button>
                </div>
            </div>
        </div>
    </div>
</div>