<script>
    import RecipeForm from './form/RecipeForm.svelte';
    import axios from 'axios';

    export let recipe;

    function submit() {
        axios.patch('http://localhost:5000/api/recipes/', recipe).then(console.log).catch(console.log)
    }
</script>

<section>
    <h1>Edit recipe {recipe.title}</h1>
    <RecipeForm onSubmit={submit} {recipe} />
</section>