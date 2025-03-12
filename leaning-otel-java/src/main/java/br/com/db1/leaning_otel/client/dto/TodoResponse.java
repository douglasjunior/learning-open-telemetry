package br.com.db1.leaning_otel.client.dto;

import java.util.Objects;

public class TodoResponse {

    private Long id;
    private String title;

    public TodoResponse(Long id, String title) {
        this.id = id;
        this.title = title;
    }

    public TodoResponse() {
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (!(o instanceof TodoResponse that)) return false;
        return Objects.equals(id, that.id) && Objects.equals(title, that.title);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id, title);
    }
}
