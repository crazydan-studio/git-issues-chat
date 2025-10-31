<script lang="ts">
    import { onMount } from 'svelte';
    import { appUser, showNotification, showUserProfileDialog, showUserPasswordDialog } from '../lib/store.js';
    import { clearAllStates } from '../lib/store.js';
    import Header from './Header.svelte';
    import Footer from './Footer.svelte';
    import ChatRoomPanel from './ChatRoomPanel.svelte';
    import MessagePanel from './MessagePanel.svelte';
    import UserProfileDialog from './UserProfileDialog.svelte';
    import UserPasswordDialog from './UserPasswordDialog.svelte';

    function lockApp() {
        // We'll implement this by changing the page state in App.svelte
        // For now, we'll just show a notification
        showNotification('info', 'App locked - this would switch to lock page in a full implementation');
    }

    function logout() {
        clearAllStates();
    }
</script>

<div class="flex flex-col h-full">
    <Header 
        on:lockApp={lockApp}
        on:logout={logout}
    />
    
    <div class="flex flex-1 overflow-hidden">
        <ChatRoomPanel class="w-1/6" />
        <MessagePanel class="w-5/6" />
    </div>
    
    <Footer />
    
    <UserProfileDialog />
    <UserPasswordDialog />
</div>