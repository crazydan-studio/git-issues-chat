<script lang="ts">
    import { appUser, showUserProfileDialog, showUserPasswordDialog } from '../lib/store.js';
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher();

    let showUserMenu = false;

    function toggleUserMenu() {
        showUserMenu = !showUserMenu;
    }

    function closeUserMenu() {
        showUserMenu = false;
    }

    function openProfile() {
        showUserProfileDialog.set(true);
        closeUserMenu();
    }

    function openPassword() {
        showUserPasswordDialog.set(true);
        closeUserMenu();
    }

    function lockApp() {
        dispatch('lockApp');
        closeUserMenu();
    }

    function logout() {
        dispatch('logout');
        closeUserMenu();
    }

    // Close menu when clicking outside
    function handleClickOutside(event: MouseEvent) {
        if (showUserMenu && event.target instanceof Element && !event.target.closest('.user-menu')) {
            closeUserMenu();
        }
    }

    // Add event listener for clicks outside
    if (typeof document !== 'undefined') {
        document.addEventListener('click', handleClickOutside);
    }
</script>

<header class="bg-white border-b border-gray-200 px-4 py-3 flex items-center justify-between">
    <div class="flex items-center">
        <div class="bg-gray-200 border-2 border-dashed rounded-xl w-8 h-8 flex items-center justify-center mr-3">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 10h.01M12 10h.01M16 10h.01M9 16H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-5l-5 5v-5z" />
            </svg>
        </div>
        <h1 class="text-xl font-bold text-gray-800">Git Issues Chat</h1>
    </div>

    <div class="relative user-menu">
        <button 
            onclick={toggleUserMenu}
            class="flex items-center space-x-2 focus:outline-none"
            aria-label="User menu"
        >
            <div class="bg-gray-200 border-2 border-dashed rounded-xl w-8 h-8 flex items-center justify-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                </svg>
            </div>
            <span class="text-gray-700 font-medium">{$appUser?.displayName || 'User'}</span>
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-500" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd" />
            </svg>
        </button>

        {#if showUserMenu}
            <div class="absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-1 z-50 border border-gray-200">
                <button 
                    onclick={openProfile}
                    class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                >
                    Profile
                </button>
                <button 
                    onclick={openPassword}
                    class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                >
                    Change Password
                </button>
                <button 
                    onclick={() => { closeUserMenu(); /* TODO: Implement action log */ }}
                    class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                >
                    Action Log
                </button>
                <button 
                    onclick={lockApp}
                    class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                >
                    Lock App
                </button>
                <button 
                    onclick={logout}
                    class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                >
                    Logout
                </button>
            </div>
        {/if}
    </div>
</header>