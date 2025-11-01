<script lang="ts">
    import { onMount } from 'svelte';
    import { appUser, showNotification, showUserProfileDialog, showUserPasswordDialog } from '../../lib/store.js';
    import { clearAllStates } from '../../lib/store.js';
    import AppMainPageHeader from './AppMainPageHeader.svelte';
    import AppMainPageFooter from './AppMainPageFooter.svelte';
    import ChatPanel from '../chat/ChatPanel.svelte';
    import UserProfileDialog from '../user/UserProfileDialog.svelte';
    import UserPasswordDialog from '../user/UserPasswordDialog.svelte';

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
    <AppMainPageHeader 
        on:lockApp={lockApp}
        on:logout={logout}
    />
    
    <ChatPanel />
    
    <AppMainPageFooter />
    
    <UserProfileDialog />
    <UserPasswordDialog />
</div>