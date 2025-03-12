package br.com.db1.leaning_otel.service;

import br.com.db1.leaning_otel.client.PlaceholderClient;
import br.com.db1.leaning_otel.exceptions.NotFoundException;
import br.com.db1.leaning_otel.model.TodoModel;
import br.com.db1.leaning_otel.repository.TodoRepository;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

@Service
public class TodoService {
    private final TodoRepository todoRepository;
    private final PlaceholderClient placeholderClient;

    public TodoService(TodoRepository todoRepository, PlaceholderClient placeholderClient) {
        this.todoRepository = todoRepository;
        this.placeholderClient = placeholderClient;
    }

    public TodoModel create(String description) {
        return todoRepository.save(new TodoModel(null, description));
    }

    public void delete(UUID id) {
        todoRepository.deleteById(id);
    }

    public TodoModel findById(UUID id) {
        return todoRepository.findById(id).orElseThrow(NotFoundException::new);
    }

    public List<Object> findAll() {
        var localTodos = todoRepository.findAll();
        var remoteTodos = placeholderClient.requestAll().getBody();

        var response = new ArrayList<>();
        response.addAll(localTodos);
        response.addAll(remoteTodos);

        return response;
    }
}
