<script lang="ts">
    import GitRepoPanel from '../repo/GitRepoPanel.svelte';
    import GitIssuePanel from '../issue/GitIssuePanel.svelte';
</script>

<div class="flex flex-1 overflow-hidden">
    <GitRepoPanel class="w-1/6" />
    <GitIssuePanel class="w-5/6" />
</div>