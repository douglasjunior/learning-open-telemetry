using LearningOtelDotnet.Model;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
// Learn more about configuring OpenAPI at https://aka.ms/aspnet/openapi
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

using var db = new AppDbContext();

app.MapGet("/todos", async () =>
{
    var todos = await db.Todos.ToListAsync();
    return todos;
})
.WithName("ListTodos");

app.Run();
