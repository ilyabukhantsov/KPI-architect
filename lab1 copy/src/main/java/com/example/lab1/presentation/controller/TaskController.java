package com.example.lab1.presentation.controller;

import com.example.lab1.application.service.TaskService;
import com.example.lab1.domain.model.Task;
import com.example.lab1.presentation.dto.TaskRequestDTO;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/api/tasks")
public class TaskController {

    private final TaskService taskService;

    // Dependency Injection через конструктор
    public TaskController(TaskService taskService) {
        this.taskService = taskService;
    }

    @GetMapping
    public List<Task> getAllTasks() {
        return taskService.getAllTasks();
    }

    @PostMapping
    public ResponseEntity<Task> createTask(@RequestBody TaskRequestDTO request) {
        Task createdTask = taskService.createTask(
            request.getName(), 
            request.getDescription(), 
            request.getDueDate()
        );
        return new ResponseEntity<>(createdTask, HttpStatus.CREATED);
    }

    @DeleteMapping("/{id}")
    @ResponseStatus(HttpStatus.NO_CONTENT)
    public void deleteTask(@PathVariable Long id) {
        taskService.deleteTask(id);
    }
}