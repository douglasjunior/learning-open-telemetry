
using LearningOtelDotnet.Model;

namespace LearningOtelDotnet.Service;

public class TodoService(AppDbContext dbContext)
{
  public List<TodoModel> GetAllTodos()
  {
    return dbContext.Todos.ToList();
  }

  public TodoModel GetTodoById(int id)
  {
    return dbContext.Todos.FirstOrDefault(t => t.TodoId == id);
  }

  public void AddTodo(TodoModel todo)
  {
    dbContext.Todos.Add(todo);
    dbContext.SaveChanges();
  }


  public void DeleteTodoById(int id)
  {
    var todo = dbContext.Todos.FirstOrDefault(t => t.TodoId == id);
    if (todo != null)
    {
      dbContext.Todos.Remove(todo);
      dbContext.SaveChanges();
    }
  }
}
