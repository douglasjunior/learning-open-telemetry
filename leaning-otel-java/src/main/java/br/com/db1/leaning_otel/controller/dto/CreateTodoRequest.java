package br.com.db1.leaning_otel.controller.dto;

import jakarta.validation.constraints.NotNull;
import org.hibernate.validator.constraints.Length;

import java.util.Objects;

public class CreateTodoRequest {

    @NotNull
    @Length(min = 1, max = 255)
    private String description;

    public CreateTodoRequest(String description) {
        this.description = description;
    }

    public CreateTodoRequest() {
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
        if (!(o instanceof CreateTodoRequest that)) return false;
        return Objects.equals(description, that.description);
    }

    @Override
    public int hashCode() {
        return Objects.hashCode(description);
    }
}
