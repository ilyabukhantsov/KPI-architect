package com.example.lab1.service;

import com.example.lab1.model.Task;
import com.example.lab1.repository.TaskRepository;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class TaskService {
    private final TaskRepository repository;

    public TaskService(TaskRepository repository){
        this.repository = repository;
    }

    public List<Task> getAll(){
        return repository.findAll();
    }

    public Task create(Task task){
        return repository.save(task);
    }

    public Task getById(Long id){
        return repository.findById(id).orElse(null);
    }

    public void delete(Long id){
        repository.deleteById(id);
    }

    public Task update(Long id, Task updatedTask) {
    return repository.findById(id)
    .map(task -> {
            task.setName(updatedTask.getName());
            task.setDescribe(updatedTask.getDescribe());
            task.setData(updatedTask.getData());
            return repository.save(task);
        })
        .orElse(null);
    }
}
