using System.ComponentModel.DataAnnotations;

namespace LearningOtelDotnet.Model;

public class TodoModel
{
    public TodoModel(string Description)
    {
        this.Description = Description;
    }

    [Key]
    public int TodoId { get; set; }
    public string Description { get; set; }

}
