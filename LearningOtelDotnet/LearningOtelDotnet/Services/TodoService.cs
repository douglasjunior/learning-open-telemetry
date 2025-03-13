
using LearningOtelDotnet.Controller.Dto;
using LearningOtelDotnet.Models;
using Microsoft.EntityFrameworkCore;
using NuGet.Protocol;

namespace LearningOtelDotnet.Services;

public class TodoService
{
    private readonly AppDbContext _context;
    private readonly ILogger<TodoService> _logger;

    public TodoService(AppDbContext context, ILogger<TodoService> logger)
    {
        _context = context;
        _logger = logger;
    }

    public async Task<IEnumerable<TodoModel>> GetTodos()
    {
        _logger.LogInformation("Get all todos");
        var todos = await _context.Todos.ToListAsync();
        return todos;
    }

    public async Task<TodoModel?> GetTodoModel(long id)
    {
        _logger.LogInformation("Get todo by id: {0}", id);
        var todoModel = await _context.Todos.FindAsync(id);
        return todoModel;
    }

    public async Task<TodoModel> PostTodoModel(CreateTodoRequest todoRequest)
    {
        _logger.LogInformation("Post todo: {0}", todoRequest.ToJson());
        var todo = (await _context.Todos.AddAsync(new TodoModel {
            Description = todoRequest.Description
        })).Entity;
        await _context.SaveChangesAsync();
        return todo;
    }

    public async Task DeleteTodoModel(long id)
    {
        _logger.LogInformation("Delete todo by id: {0}", id);
        var todoModel = await _context.Todos.FindAsync(id);
        if (todoModel == null)
        {
            return;
        }
        _context.Todos.Remove(todoModel);
        await _context.SaveChangesAsync();
    }
}
