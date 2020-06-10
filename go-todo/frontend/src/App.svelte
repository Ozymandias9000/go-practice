<script>
  import { onMount } from "svelte";
  import { http } from "./utils/http";
  import * as api from "./api/api";

  let todos = [];
  let newTodoValue = "";

  onMount(async () => {
    try {
      const { data } = await http.get("/todos");

      todos = data.todos;
    } catch (e) {
      console.error(e);
    }
  });

  const handleNewTodo = async () => {
    try {
      const todo = await api.addNewTodo(newTodoValue);
      console.log(todos, todo);
      todos = [...todos, todo];
    } catch (e) {
      console.error(e); // eslint-disable-line
    }
  };

  const handleUpdateTodo = todo => async () => {
    try {
      const updatedTodo = await api.updateTodo(todo);

      todos = todos.map(t => {
        if (t.id === todo.id) {
          return updatedTodo;
        }

        return t;
      });
      console.log(todos, updatedTodo);
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
      <button
        on:click={handleUpdateTodo({ ...todo, completed: !todo.completed })}>
        {#if todo.completed}√{:else}x{/if}
      </button>
    </div>
  {/each}
</div>
