<script>
  import { onMount } from "svelte";
  import { http } from "./utils/http";
  import { addNewTodo } from "./api/api";

  let todos = [];
  let newTodoValue = "";

  onMount(async () => {
    try {
      const jwt =
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTE3MzgxNzIsImlkIjoxOX0.d4iEl1F7yxRhOp8pYooAnNpwHGklFtkTm6bsAMoOMEg";

      const { data } = await http.get("/todos");

      todos = data.todos;
    } catch (e) {
      console.error(e);
    }
  });

  const handleNewTodo = async () => {
    try {
      const todo = await addNewTodo(newTodoValue);
      console.log(todos, todo);
      todos = [...todos, todo];
    } catch (e) {
      console.error(e); // eslint-disable-line
    }
  };
</script>

<style>

</style>

<div class="App">
  <input type="text" bind:value={newTodoValue} />
  <button on:click={handleNewTodo} type="submit">Make Todo</button>

  {#each todos as todo}
    <div>
      {todo.body}
      <button>
        {#if todo.completed}√{:else}x{/if}
      </button>
    </div>
  {/each}
</div>
