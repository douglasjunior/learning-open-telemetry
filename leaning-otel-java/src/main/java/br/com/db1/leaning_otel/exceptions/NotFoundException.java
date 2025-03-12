package br.com.db1.leaning_otel.exceptions;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ResponseStatus;

@ResponseStatus(value = HttpStatus.NOT_FOUND, reason = "No such record")
public class NotFoundException extends RuntimeException {
}