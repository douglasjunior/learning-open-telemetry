using System.ComponentModel.DataAnnotations;

namespace LearningOtelDotnet.Controller.Dto;

public class CreateTodoRequest
{

    [Required]
    [StringLength(255)]
    public required string Description { get; set; }

}
