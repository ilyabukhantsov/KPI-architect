package com.example.lab1.domain.factory;

import com.example.lab1.domain.model.Task;
import java.time.LocalDate;


public class TaskFactory {
    public static Task createNewTask(String name, String description, LocalDate dueDate) {
        // Тут можна додати складні інваріанти (наприклад, дата не в минулому)
        return new Task(null, name, description, dueDate);
    }
}