<script>
    import axios from 'axios';

    async function getRecipes() {
        const resp = await axios.get('http://localhost:5000/api/recipes').catch(console.log)
        return resp.data;
    }

    const recipesReq = getRecipes()
</script>

<section>
    RECIPES

    {#await recipesReq}
        <span>Loading recipes...</span>
    {:then recipes}
        {#each recipes as recipe}
            <div>
            <h4>{recipe.Title}</h4>
            <p>{recipe.Content}</p>
            </div>
        {/each}
    {:catch error}
        <span>Error: Problem loading recipes - {error}</span>
    {/await}
</section>