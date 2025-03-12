using LearningOtelDotnet.Controller.Dto;
using Microsoft.AspNetCore.Mvc;
using LearningOtelDotnet.Service;
using LearningOtelDotnet.Model;

namespace LearningOtelDotnet.Controller;

[ApiController]
[Route("todos")]
public class TodoController(TodoService todoService) : ControllerBase
{
  [HttpPost]
  public IActionResult CreateTodo([FromBody] CreateTodoRequest createTodoRequest)
  {
    if (!ModelState.IsValid)
    {
      return BadRequest(ModelState);
    }

    var todo = todoService.AddTodo( new TodoModel(Description: createTodoRequest.Description));

    return todo;
  }

  [HttpGet("{id}")]
  public IActionResult GetTodo(int id)
  {
    // Get the todo by id
    var todo = todoService.GetTodoById(id);

    // Check if the todo exists
    if (todo == null)
    {
      return NotFound();
    }

    // Return the todo
    return Ok(todo);
  }
}


