package com.example.lab1.infrastructure.mapper;

import com.example.lab1.domain.model.Task;
import com.example.lab1.infrastructure.entity.TaskEntity;
import org.springframework.stereotype.Component;

@Component
public class TaskMapper {
    public TaskEntity toEntity(Task domain) {
        return new TaskEntity(domain.getId(), domain.getName(), domain.getDescription(), domain.getDueDate());
    }

    public Task toDomain(TaskEntity entity) {
        return new Task(entity.getId(), entity.getName(), entity.getDescription(), entity.getDueDate());
    }
}