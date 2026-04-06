package com.example.lab1.domain;

import com.example.lab1.domain.model.Task;
import com.example.lab1.domain.model.DomainException;
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

class TaskTest {
    @Test
    void shouldThrowExceptionWhenNameIsInvalid() {
        assertThrows(DomainException.class, () -> new Task(null, "", "Desc", null));
    }

    @Test
    void shouldCreateTaskSuccessfully() {
        Task task = new Task(1L, "Test", "Desc", null);
        assertEquals("Test", task.getName());
    }
}