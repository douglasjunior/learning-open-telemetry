using LearningOtelDotnet.Services;

namespace LearningOtelDotnet.Configs;

public static class ServicesConfig
{
    public static void AddServices(this IServiceCollection services)
    {
        services.AddScoped<TodoService>();
    }
}
