package br.com.db1.leaning_otel.repository;

import br.com.db1.leaning_otel.model.TodoModel;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public interface TodoRepository extends JpaRepository<TodoModel, UUID> {
}
