import { http } from "../utils/http";

export const addNewTodo = async (text) => {
  try {
    const {
      data: { todo },
    } = await http.post("/todos", { body: text });

    return todo;
  } catch (e) {
    throw e;
  }
};

export const updateTodo = async (updatedTodo) => {
  try {
    const {
      data: { todo },
    } = await http.patch(`/todos/${updatedTodo.id}`, updatedTodo);

    return todo;
  } catch (e) {
    throw e;
  }
};
