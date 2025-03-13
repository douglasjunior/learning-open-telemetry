
using Microsoft.EntityFrameworkCore;

namespace LearningOtelDotnet.Models;

public class AppDbContext : DbContext
{
    public DbSet<TodoModel> Todos { get; set; }

    private ILogger<AppDbContext> _logger;
    public string DbPath { get; }

    public AppDbContext(ILogger<AppDbContext> logger)
    {
        _logger = logger;
        var path = Path.GetDirectoryName(AppDomain.CurrentDomain.BaseDirectory);
        DbPath = Path.Join(path, "todolist.db");
        _logger.LogInformation("DbPath: {0}", DbPath);
    }

    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        => optionsBuilder.UseSqlite($"Data Source={DbPath}");
}
