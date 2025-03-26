package br.com.db1.leaning_otel.service;

import br.com.db1.leaning_otel.client.PlaceholderClient;
import br.com.db1.leaning_otel.exceptions.NotFoundException;
import br.com.db1.leaning_otel.model.TodoModel;
import br.com.db1.leaning_otel.repository.TodoRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

@Service
public class TodoService {
    private static final Logger log = LoggerFactory.getLogger(TodoService.class);

    private final TodoRepository todoRepository;
    private final PlaceholderClient placeholderClient;

    public TodoService(TodoRepository todoRepository, PlaceholderClient placeholderClient) {
        this.todoRepository = todoRepository;
        this.placeholderClient = placeholderClient;
    }

    public TodoModel create(String description) {
        log.info("Creating todo with description: {}", description);
        return todoRepository.save(new TodoModel(null, description));
    }

    public void delete(UUID id) {
        log.info("Deleting todo with id: {}", id);
        todoRepository.deleteById(id);
    }

    public TodoModel findById(UUID id) {
        log.info("Finding todo with id: {}", id);
        return todoRepository.findById(id).orElseThrow(NotFoundException::new);
    }

    public List<Object> findAll() {
        log.info("Finding all todos");
        var localTodos = todoRepository.findAll();
        var remoteTodos = placeholderClient.requestAll().getBody();

        var response = new ArrayList<>();
        response.addAll(localTodos);
        response.addAll(remoteTodos);

        return response;
    }

    @SuppressWarnings("ConstantConditions")
    public void error() {
        log.info("Throwing exception");
        String nullString = null;
        var ignored = nullString.split("");
    }
}
