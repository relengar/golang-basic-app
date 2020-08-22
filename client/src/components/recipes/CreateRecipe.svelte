<script>
    import RecipeForm from './form/RecipeForm.svelte';
    import axios from 'axios';

    const recipe = {}

    function submit() {
        axios.post('http://localhost:5000/api/recipes/', recipe).then(console.log).catch(console.log)
    }
</script>

<section>
    <h1>Create new recipe</h1>
    <RecipeForm onSubmit={submit} {recipe} />

</section>