
using LearningOtelDotnet.Client;
using LearningOtelDotnet.Controller.Dto;
using LearningOtelDotnet.Models;
using Microsoft.EntityFrameworkCore;
using NuGet.Protocol;

namespace LearningOtelDotnet.Services;

public class TodoService
{
    private readonly AppDbContext _context;
    private readonly ILogger<TodoService> _logger;
    private readonly PlaceholderClient _placeholderClient;

    public TodoService(AppDbContext context, ILogger<TodoService> logger, PlaceholderClient placeholderClient)
    {
        _context = context;
        _logger = logger;
        _placeholderClient = placeholderClient;
    }

    public async Task<IEnumerable<object>> GetTodos()
    {
        _logger.LogInformation("Get all todos");

        var list = new List<object>();
        var todos = await _context.Todos.ToListAsync();
        var externalTodos = await _placeholderClient.RequestTodos();

        list.AddRange(todos);
        list.AddRange(externalTodos);

        return list;
    }

    public async Task<TodoModel?> GetTodoModel(long id)
    {
        _logger.LogInformation("Get todo by id: {0}", id);
        var todo = await _context.Todos.FindAsync(id);
        return todo;
    }

    public async Task<TodoModel> PostTodoModel(CreateTodoRequest todoRequest)
    {
        _logger.LogInformation("Post todo: {0}", todoRequest.ToJson());
        var todo = (await _context.Todos.AddAsync(new TodoModel
        {
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

    internal async Task Error()
    {
        _logger.LogInformation("Throwing exception");
        string? text = null;
        _ = text!.Split(" ");
    }
}
