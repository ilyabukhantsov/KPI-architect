package com.example.lab1.domain.model;

import java.time.LocalDate;


public class Task {
    private final Long id;
    private String name;
    private String description;
    private LocalDate dueDate;

    public Task(Long id, String name, String description, LocalDate dueDate) {
        this.id = id;
        setName(name);
        this.description = description;
        this.dueDate = dueDate;
    }

    public void setName(String name) {
        if (name == null || name.isBlank()) {
            throw new DomainException("Назва не може бути порожньою");
        }
        this.name = name;
    }

    // Геттери
    public Long getId() { return id; }
    public String getName() { return name; }
    public String getDescription() { return description; }
    public LocalDate getDueDate() { return dueDate; }
}