using Microsoft.AspNetCore.Mvc;
using LearningOtelDotnet.Services;
using LearningOtelDotnet.Models;
using LearningOtelDotnet.Controller.Dto;

namespace LearningOtelDotnet.Controllers;

[ApiController]
[Route("[controller]")]
public class TodoController : ControllerBase
{
    private readonly TodoService _service;

    public TodoController(TodoService service)
    {
        _service = service;
    }

    [HttpGet]
    public async Task<ActionResult<IEnumerable<object>>> GetTodos()
    {
        return Ok(await _service.GetTodos());
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<TodoModel>> GetTodoModel(long id)
    {
        var todoModel = await _service.GetTodoModel(id);
        if (todoModel == null)
        {
            return NotFound();
        }
        return Ok(todoModel);
    }

    [HttpPost]
    public async Task<ActionResult<TodoModel>> PostTodoModel(CreateTodoRequest todoRequest)
    {
        var todo = await _service.PostTodoModel(todoRequest);
        return CreatedAtAction(nameof(GetTodoModel), new { id = todo.TodoId }, todo);
    }

    [HttpDelete("{id}")]
    public async Task<IActionResult> DeleteTodoModel(long id)
    {
        await _service.DeleteTodoModel(id);
        return NoContent();
    }

    [HttpGet("Error")]
    public async Task<IActionResult> Error()
    {
        await _service.Error();
        return Ok();
    }
}
