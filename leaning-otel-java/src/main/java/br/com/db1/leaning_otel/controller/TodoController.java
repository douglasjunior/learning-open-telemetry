package br.com.db1.leaning_otel.controller;

import br.com.db1.leaning_otel.controller.dto.CreateTodoRequest;
import br.com.db1.leaning_otel.model.TodoModel;
import br.com.db1.leaning_otel.service.TodoService;
import jakarta.validation.Valid;
import jakarta.validation.constraints.NotNull;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;
import java.util.UUID;

@RestController()
@RequestMapping("todos")
@Valid
public class TodoController {

    private final TodoService todoService;

    public TodoController(TodoService todoService) {
        this.todoService = todoService;
    }

    @GetMapping("{id}")
    public ResponseEntity<TodoModel> get(@NotNull @PathVariable("id") UUID id) {
        var todo = todoService.findById(id);
        return ResponseEntity.ok(todo);
    }

    @GetMapping
    public ResponseEntity<List<Object>> getAll() {
        var todos = todoService.findAll();
        return ResponseEntity.ok(todos);
    }

    @DeleteMapping("{id}")
    public ResponseEntity<Void> delete(@NotNull @PathVariable("id") UUID id) {
        todoService.delete(id);
        return ResponseEntity.noContent().build();
    }

    @PostMapping
    public ResponseEntity<TodoModel> create(@RequestBody CreateTodoRequest todoRequest) {
        var todo = todoService.create(todoRequest.getDescription());
        return ResponseEntity.created(null).body(todo);
    }

}
