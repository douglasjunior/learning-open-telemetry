
using Microsoft.EntityFrameworkCore;

namespace LearningOtelDotnet.Models;

public class AppDbContext : DbContext
{
    public DbSet<TodoModel> Todos { get; set; }

    public string DbPath { get; }

    public AppDbContext()
    {
        var path = Path.GetDirectoryName(AppDomain.CurrentDomain.BaseDirectory);
        DbPath = Path.Join(path, "todolist.db");
        Console.WriteLine($"DbPath: {DbPath}");
    }

    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        => optionsBuilder.UseSqlite($"Data Source={DbPath}");
}
