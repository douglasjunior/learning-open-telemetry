package br.com.db1.leaning_otel.model;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

import java.util.Objects;
import java.util.UUID;

@Entity
@Table(name = "todo")
public class TodoModel {

    @GeneratedValue(generator = "UUID")
    @Id
    private UUID id;

    @Column(length = 255, nullable = false)
    private String description;

    public TodoModel() {
    }

    public TodoModel(UUID id, String description) {
        this.id = id;
        this.description = description;
    }

    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (!(o instanceof TodoModel todoModel)) return false;
        return Objects.equals(id, todoModel.id) && Objects.equals(description, todoModel.description);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, description);
    }
}
