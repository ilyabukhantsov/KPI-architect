package com.example.lab1.controller;

import com.example.lab1.model.Task;
import com.example.lab1.service.TaskService;
import org.springframework.web.bind.annotation.*;

import java.util.List;


@RestController
@RequestMapping("/tasks")
public class TaskController {
    private final TaskService service;

    public TaskController(TaskService service){
        this.service = service;
    }

    @GetMapping
    public List<Task> getAll(){
        return service.getAll();
    }

    @PostMapping
    public Task create(@RequestBody Task task){
        return service.create(task);
    }

    @GetMapping("/{id}")
    public Task getById(@PathVariable Long id){
        return service.getById(id);
    }



    @DeleteMapping("/{id}")
    public void delete(@PathVariable Long id){
        service.delete(id);
    }

    @PutMapping("/{id}")
    public Task updateTask(@PathVariable Long id, @RequestBody Task task) {
        return service.update(id, task);
    }
}
