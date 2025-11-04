import { toast } from 'svoast';
import CustomToast from './CustomToast.svelte';

// Helper functions for different notification types using Svoast
export function info(message: string, duration: number = 3000) {
  toast.info(message, {
    component: [CustomToast, {}],
    duration,
    closable: true,
    infinite: false
  });
}

export function success(message: string, duration: number = 3000) {
  toast.success(message, {
    component: [CustomToast, {}],
    duration,
    closable: true,
    infinite: false
  });
}

export function error(message: string, duration: number = 0) {
  // Error notifications don't auto-close by default
  toast.error(message, {
    component: [CustomToast, {}],
    duration: 0,
    closable: true,
    infinite: true
  });
}

export function warning(message: string, duration: number = 3000) {
  toast.warning(message, {
    component: [CustomToast, {}],
    duration,
    closable: true,
    infinite: false
  });
}

// Notify object to export all functions
export const Notify = {
  info,
  success,
  error,
  warning
};

export { default as Notification } from './Notification.svelte';

