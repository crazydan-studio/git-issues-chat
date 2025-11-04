<script lang="ts">
  import { Toast } from 'svoast';
  import type { ToastComponent, ToastType } from 'svoast';

  // Accept individual properties as specified in the requirements
  let { 
    id,
    type,
    message,
    closable,
    duration,
    infinite,
    rich
  }: { 
    id: number;
    type: ToastType;
    message: string;
    closable: boolean;
    duration: number;
    infinite: boolean;
    rich: boolean;
  } = $props();

  // Create a toast object to pass to the Toast component
  // Exclude duration as per requirements
  let toast: ToastComponent = $state({
    id,
    type,
    message,
    closable,
    infinite,
    rich,
    duration
  });
</script>

<div class={["customToast w-fit", type, !infinite && 'progress']} style="--duration: {duration}ms">
  <Toast {toast} />
</div>

<style>
  .customToast {
    --svoast-radius: 4px;
    --svoast-bar-width: 4px;

    &  :global(.svoast) {
      --svoast-bg: #ffffff;
      --svoast-text: #333333;
      --svoast-shadow: 0 2px 7px hsl(0 0% 0% / 0.1);
    }

    &.progress {
      & :global(.svoast-bar) {
        top: unset;
        bottom: 0;
        border-top-left-radius: var(--svoast-radius);
        border-top-right-radius: var(--svoast-radius);
        animation: timeBar var(--duration, 3000ms) linear forwards;
      }
    }

    &.info {
      --svoast-info-colour: #6b7280;
      --svoast-colour: var(--svoast-info-colour);
    }
    
    &.attention {
      --svoast-attention-colour: #0ea5e9;
      --svoast-colour: var(--svoast-attention-colour);
    }
    
    &.success {
      --svoast-success-colour: #22c55e;
      --svoast-colour: var(--svoast-success-colour);
    }
    
    &.warning {
      --svoast-warning-colour: #f97316;
      --svoast-colour: var(--svoast-warning-colour);
    }
    
    &.error {
      --svoast-error-colour: #ef4444;
      --svoast-colour: var(--svoast-error-colour);
    }
  }
  
  @keyframes timeBar {
    from { height: 100%; }
    to { height: 0%; display: none; }
  }
</style>
