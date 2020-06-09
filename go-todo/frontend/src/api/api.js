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
