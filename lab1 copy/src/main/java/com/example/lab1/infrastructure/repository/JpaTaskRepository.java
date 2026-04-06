package com.example.lab1.infrastructure.repository;

import com.example.lab1.infrastructure.entity.TaskEntity;
import org.springframework.data.jpa.repository.JpaRepository;

public interface JpaTaskRepository extends JpaRepository<TaskEntity, Long> {
}