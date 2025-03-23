using LearningOtelDotnet.Client;
using LearningOtelDotnet.Services;

namespace LearningOtelDotnet.Configs;

public static class ServicesConfig
{
    public static void AddHttpClients(this IServiceCollection services, ConfigurationManager configuration)
    {
        services.AddHttpClient<PlaceholderClient>(client =>
        {
            client.BaseAddress = new Uri(configuration.GetValue<string>("Placeholder:BaseUrl")!);
        });
    }

    public static void AddServices(this IServiceCollection services)
    {
        services.AddScoped<TodoService>();
    }
}
