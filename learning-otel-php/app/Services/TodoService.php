<?php
namespace App\Services;

use App\Models\Todo;
use Log;

class TodoService
{
    public function getAllTodos()
    {
        Log::info('Getting all todos');
        return Todo::all();
    }

    public function getTodoById($id)
    {
        Log::info('Getting todo by id', ['id' => $id]);
        return Todo::findOrFail($id);
    }

    public function createTodo(array $data)
    {
        Log::info('Creating todo', ['data' => $data]);
        return Todo::create($data);
    }

    public function deleteTodo($id)
    {
        Log::info('Deleting todo', ['id' => $id]);
        $todo = Todo::findOrFail($id);
        $todo->delete();
    }

    public function error()
    {
        Log::info('Cause error');
        $text = null;
        $text->test();
    }

}
