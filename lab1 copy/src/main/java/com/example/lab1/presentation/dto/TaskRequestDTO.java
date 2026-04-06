package com.example.lab1.presentation.dto;

import lombok.Data;
import java.time.LocalDate;

@Data
public class TaskRequestDTO {
    private String name;
    private String description;
    private LocalDate dueDate;
}