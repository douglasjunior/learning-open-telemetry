<?php
namespace App\Services;

use App\Models\Todo;

class TodoService
{
    public function getAllTodos()
    {
        return Todo::all();
    }

    public function getTodoById($id)
    {
        return Todo::findOrFail($id);
    }

    public function createTodo(array $data)
    {
        return Todo::create($data);
    }

    public function deleteTodo($id)
    {
        $todo = Todo::findOrFail($id);
        $todo->delete();
    }

}
