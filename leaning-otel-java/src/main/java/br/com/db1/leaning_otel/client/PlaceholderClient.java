package br.com.db1.leaning_otel.client;

import br.com.db1.leaning_otel.client.dto.TodoResponse;
import org.springframework.cloud.openfeign.FeignClient;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;

import java.util.List;

@FeignClient(
        name = "PlaceholderClient",
        url = "${placeholder.base-url}"
)
public interface PlaceholderClient {

    @GetMapping("/todos")
    ResponseEntity<List<TodoResponse>> requestAll();

}
