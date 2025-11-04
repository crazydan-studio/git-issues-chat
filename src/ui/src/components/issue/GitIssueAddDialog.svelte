<script lang="ts">
    import { showAddIssueDialog } from '../../lib/store';
    import { saveGitIssue } from '../../lib/bridge';
    import { Notify } from '../../lib/components/Notification';
    import Dialog from '../../lib/components/Dialog.svelte';
    import { MarkdownEditor } from '../../lib/components/Markdown';

    let title = $state('');
    let content = $state('');
    let isSaving = $state(false);

    function closeDialog() {
        showAddIssueDialog.set(false);
        // Reset form
        title = '';
        content = '';
    }

    async function saveIssue() {
        if (!title) {
            Notify.warning('Issue title is required');
            return;
        }

        isSaving = true;
        const result = await saveGitIssue({
            title,
            content
        });

        if (result.success) {
            Notify.success('Issue created successfully');
            closeDialog();
        } else {
            Notify.error(result.msg || 'Failed to create issue');
        }
        isSaving = false;
    }
</script>

{#if $showAddIssueDialog}
    <Dialog 
        title="Create New Issue"
        onClose={closeDialog}
        onConfirm={saveIssue}
        class="w-4/5"
    >
        <div class="space-y-4">
            <div>
                <label for="issueTitle" class="block text-sm font-medium text-gray-700 mb-1">Title *</label>
                <input
                    id="issueTitle"
                    type="text"
                    bind:value={title}
                    class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                    placeholder="Enter issue title"
                />
            </div>

            <div>
                <label for="issueContent" class="block text-sm font-medium text-gray-700 mb-1">Content</label>
                <div class="w-full">
                    <MarkdownEditor
                        bind:value={content}
                    />
                </div>
            </div>
        </div>
    </Dialog>
{/if}
