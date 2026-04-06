package com.example.lab1.infrastructure.entity;

import jakarta.persistence.*;
import lombok.*; // Обов'язково!
import java.time.LocalDate;

@Entity
@Table(name = "tasks")
@Getter @Setter 
@NoArgsConstructor 
@AllArgsConstructor
public class TaskEntity {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    private String name;
    private String description;
    private LocalDate dueDate;
}