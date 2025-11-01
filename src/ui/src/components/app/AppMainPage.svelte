<script lang="ts">
    import { clearAllStates } from '../../lib/store.js';
    import AppMainPageHeader from './AppMainPageHeader.svelte';
    import AppMainPageFooter from './AppMainPageFooter.svelte';
    import ChatPanel from '../chat/ChatPanel.svelte';
    import AppUserActionLogPanel from './AppUserActionLogPanel.svelte';
    import UserProfileDialog from '../user/UserProfileDialog.svelte';
    import UserPasswordDialog from '../user/UserPasswordDialog.svelte';

    let { onLockApp, onLogout }: { onLockApp: () => void; onLogout: () => void } = $props();
    let showActionLogPanel = $state(false);

    function showActionLog() {
        showActionLogPanel = true;
    }

    function hideActionLog() {
        showActionLogPanel = false;
    }
</script>

<div class="flex flex-col h-full">
    <AppMainPageHeader 
        onLockApp={onLockApp}
        onLogout={onLogout}
        onShowActionLog={showActionLog}
    />
    
    {#if showActionLogPanel}
        <AppUserActionLogPanel onBack={hideActionLog} />
    {:else}
        <ChatPanel />
    {/if}
    
    <AppMainPageFooter />
    
    <UserProfileDialog />
    <UserPasswordDialog />
</div>