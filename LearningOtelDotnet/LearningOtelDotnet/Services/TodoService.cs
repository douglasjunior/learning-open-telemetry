
using LearningOtelDotnet.Models;
using Microsoft.EntityFrameworkCore;

namespace LearningOtelDotnet.Services;

public class TodoService
{
    private readonly AppDbContext _context;

    public TodoService(AppDbContext context)
    {
        _context = context;
    }

    public async Task<IEnumerable<TodoModel>> GetTodos()
    {
        var todos = await _context.Todos.ToListAsync();
        return todos;
    }

    public async Task<TodoModel?> GetTodoModel(long id)
    {
        var todoModel = await _context.Todos.FindAsync(id);
        return todoModel;
    }

    public async Task<TodoModel> PostTodoModel(TodoModel todoModel)
    {
        var todo = (await _context.Todos.AddAsync(todoModel)).Entity;
        await _context.SaveChangesAsync();
        return todo;
    }

    public async Task DeleteTodoModel(long id)
    {
        var todoModel = await _context.Todos.FindAsync(id);
        if (todoModel == null)
        {
            return;
        }
        _context.Todos.Remove(todoModel);
        await _context.SaveChangesAsync();
    }
}
