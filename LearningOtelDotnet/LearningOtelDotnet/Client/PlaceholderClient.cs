using LearningOtelDotnet.Client.Dto;
using Newtonsoft.Json;

namespace LearningOtelDotnet.Client;

public class PlaceholderClient
{

    private readonly HttpClient _httpClient;

    public PlaceholderClient(HttpClient httpClient)
    {
        _httpClient = httpClient;
    }

    public async Task<IEnumerable<TodoResponse>> RequestTodos()
    {
        var responseString = await _httpClient.GetStringAsync("/todos");
        return JsonConvert.DeserializeObject<IEnumerable<TodoResponse>>(responseString)!;
    }
}
