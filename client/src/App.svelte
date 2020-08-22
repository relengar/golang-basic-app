<script>
	import Recipe from './components/recipes/Recipe.svelte'
	import CreateRecipe from './components/recipes/CreateRecipe.svelte'
	import EditRecipe from './components/recipes/EditRecipe.svelte'
	import { Router, Link, Route } from 'svelte-routing';

	const routes = [
		{
			path: '/',
			name: 'recipes',
			component: Recipe
		},
		{
			path: '/recipes/new',
			name: 'createRecipe',
			component: CreateRecipe
		},
		{
			path: '/recipes/edit/:id',
			name: 'recipesEdit',
			component: EditRecipe
		},
	]


	export let url = "";
</script>

<style>
	:global(body) {
		margin: 0;
		font-family: Arial, Helvetica, sans-serif;
		background-color: #F9F6F6;
	}
	.App {
		text-align: center;
	}
	.App-header {	
		color: #333;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		font-size: calc(10px + 2vmin);
	}
</style>

<div class="App">
	<!-- <Router url="{url}"> -->
		<header class="App-header">head
			<!-- <Link to={'/recipes'}>Recipes</Link> -->
			<!-- <Link to={'/recipes/new'}>Create Recipe</Link> -->
		</header>
		<!-- <Route path="/recipes" component="{Recipe}" /> -->
		<!-- {#each routes as route}
		{route.path}
            <Route path={route.path} component={route.component} />
        {/each} -->
	<!-- </Router> -->

</div>