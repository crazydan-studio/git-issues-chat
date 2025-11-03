<script lang="ts">
    import { notification } from '../lib/store';
    import { AlertCircle, CheckCircle, Info } from 'lucide-svelte';

    // Auto-close non-error notifications after 3 seconds
    $effect(() => {
        if ($notification.message && $notification.type !== 'error') {
            const timer = setTimeout(() => {
                notification.set({ message: '', type: 'info' });
            }, 3000);
            return () => clearTimeout(timer);
        }
    });
</script>

{#if $notification.message}
    <div class={`fixed z-50 ${$notification.type === 'error' ? 'top-4 right-4' : 'top-4 left-1/2 transform -translate-x-1/2'}`}>
        <div class={`px-6 py-4 rounded-md shadow-lg text-white font-medium flex items-center space-x-2 ${
            $notification.type === 'error' ? 'bg-red-500' : 
            $notification.type === 'success' ? 'bg-green-500' : 
            'bg-blue-500'
        }`}>
            <span>{$notification.message}</span>
            {#if $notification.type === 'error'}
                <AlertCircle class="h-5 w-5" />
            {:else if $notification.type === 'success'}
                <CheckCircle class="h-5 w-5" />
            {:else}
                <Info class="h-5 w-5" />
            {/if}
        </div>
    </div>
{/if}