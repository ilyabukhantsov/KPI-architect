package com.example.lab1.domain.repository;

import com.example.lab1.domain.model.Task;
import java.util.List;
import java.util.Optional;


public interface TaskRepository {
    Task save(Task task);
    List<Task> findAll();
    Optional<Task> findById(Long id);
    void deleteById(Long id);
}