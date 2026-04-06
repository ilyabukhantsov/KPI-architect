package com.example.lab1.infrastructure.repository;

import com.example.lab1.domain.model.Task;
import com.example.lab1.domain.repository.TaskRepository;
import com.example.lab1.infrastructure.mapper.TaskMapper;
import org.springframework.stereotype.Repository;
import java.util.List;
import java.util.Optional;
import java.util.stream.Collectors;

@Repository
public class TaskRepositoryImpl implements TaskRepository {
    private final JpaTaskRepository jpaRepo;
    private final TaskMapper mapper;

    public TaskRepositoryImpl(JpaTaskRepository jpaRepo, TaskMapper mapper) {
        this.jpaRepo = jpaRepo;
        this.mapper = mapper;
    }

    @Override
    public Task save(Task task) {
        var entity = jpaRepo.save(mapper.toEntity(task));
        return mapper.toDomain(entity);
    }

    @Override
    public List<Task> findAll() {
        return jpaRepo.findAll().stream().map(mapper::toDomain).collect(Collectors.toList());
    }

    @Override
    public Optional<Task> findById(Long id) {
        return jpaRepo.findById(id).map(mapper::toDomain);
    }

    @Override
    public void deleteById(Long id) { jpaRepo.deleteById(id); }
}