<script>
    import { onMount } from 'svelte';
    import { appUser } from './lib/store.js';
    import AppAuthPage from './components/app/AppAuthPage.svelte';
    import AppMainPage from './components/app/AppMainPage.svelte';
    import AppLockPage from './components/app/AppLockPage.svelte';
    import Notification from './components/Notification.svelte';

    let currentPage = 'auth'; // 'auth', 'chat', 'lock'

    $: {
        if ($appUser) {
            currentPage = currentPage === 'lock' ? 'lock' : 'chat';
        } else {
            currentPage = 'auth';
        }
    }

    // Function to handle lock app event
    function handleLockApp() {
        currentPage = 'lock';
    }

    // Function to handle logout event
    function handleLogout() {
        currentPage = 'auth';
    }
</script>

<div class="w-full h-screen">
    {#if currentPage === 'auth'}
        <AppAuthPage />
    {:else if currentPage === 'chat'}
        <AppMainPage on:lockApp={handleLockApp} on:logout={handleLogout} />
    {:else if currentPage === 'lock'}
        <AppLockPage />
    {/if}
    <Notification />
</div>