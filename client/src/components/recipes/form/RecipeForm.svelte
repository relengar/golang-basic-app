<script>
    const noSubmit = () => {};
    export let recipe
    export let onSubmit = noSubmit
</script>

<style>
    .button {
        width: 150px;
        border: 1px solid rgba(129, 88, 12, 0.37);
        border-radius: 10px;
        font-weight: 500;
        font-size: 18px;
        padding: 3px 0 3px 0;
    }
    
</style>

<form on:submit|preventDefault={onSubmit}>
    <label for="title" >Title</label>
    <input type="text" name="title" bind:value={recipe.title} />
    <label for="content">Content</label>
    <input type="text" name="content" bind:value={recipe.content} />
    {#if onSubmit !== noSubmit}
        <p><button class="button">Submit</button></p>
    {/if}
</form>