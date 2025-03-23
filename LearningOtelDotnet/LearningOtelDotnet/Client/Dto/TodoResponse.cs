using Newtonsoft.Json;

namespace LearningOtelDotnet.Client.Dto;

public class TodoResponse
{
    [JsonProperty("id")]
    public long Id { get; set; }

    [JsonProperty("title")]
    public required string Title { get; set; }
}
