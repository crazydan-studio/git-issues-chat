<script>
    import { onMount } from 'svelte';
    import { appUser } from './lib/store.js';
    import AuthPage from './components/AuthPage.svelte';
    import ChatPage from './components/ChatPage.svelte';
    import LockPage from './components/LockPage.svelte';
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
        <AuthPage />
    {:else if currentPage === 'chat'}
        <ChatPage on:lockApp={handleLockApp} on:logout={handleLogout} />
    {:else if currentPage === 'lock'}
        <LockPage />
    {/if}
    <Notification />
</div>