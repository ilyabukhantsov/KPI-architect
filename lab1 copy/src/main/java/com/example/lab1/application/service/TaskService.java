package com.example.lab1.application.service;

import com.example.lab1.domain.factory.TaskFactory;
import com.example.lab1.domain.model.Task;
import com.example.lab1.domain.repository.TaskRepository;
import org.springframework.stereotype.Service;
import java.time.LocalDate;
import java.util.List;

@Service
public class TaskService {
    private final TaskRepository repository;

    public TaskService(TaskRepository repository) {
        this.repository = repository;
    }

    public Task createTask(String name, String description, LocalDate date) {
        Task task = TaskFactory.createNewTask(name, description, date);
        return repository.save(task);
    }

    public List<Task> getAllTasks() {
        return repository.findAll();
    }

    // ДОДАЙ ЦЕЙ МЕТОД:
    public void deleteTask(Long id) {
        repository.deleteById(id);
    }
}