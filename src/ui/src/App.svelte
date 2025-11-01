<script>
    import { appUser, clearAllStates } from './lib/store.js';
    import AppAuthPage from './components/app/AppAuthPage.svelte';
    import AppMainPage from './components/app/AppMainPage.svelte';
    import AppLockDialog from './components/app/AppLockDialog.svelte';
    import Notification from './components/Notification.svelte';

    let currentPage = $state('auth'); // 'auth' or 'chat'
    let showLockDialog = $state(false);

    $effect(() => {
        if ($appUser) {
            currentPage = 'chat';
        } else {
            currentPage = 'auth';
        }
    });

    // Function to handle lock app event
    function handleLockApp() {
        showLockDialog = true;
    }

    // Function to handle logout event
    function handleLogout() {
        // Clear all application state data before logging out
        clearAllStates();
        currentPage = 'auth';
    }

    function closeLockDialog() {
        showLockDialog = false;
    }
</script>

<div class="w-full h-screen">
    {#if currentPage === 'auth'}
        <AppAuthPage />
    {:else if currentPage === 'chat'}
        <AppMainPage onLockApp={handleLockApp} onLogout={handleLogout} />
    {/if}
    {#if showLockDialog}
        <AppLockDialog onClose={closeLockDialog} />
    {/if}
    <Notification />
</div>