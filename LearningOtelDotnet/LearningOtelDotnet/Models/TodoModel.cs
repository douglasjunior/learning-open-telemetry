using System.ComponentModel.DataAnnotations;

namespace LearningOtelDotnet.Models;

public class TodoModel
{
    [Key]
    public long TodoId { get; set; }

    [Required]
    public string Description { get; set; }
}
