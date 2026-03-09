package com.example.lab1.model;

import java.time.LocalDate;

import jakarta.persistence.*;

@Entity
public class Task {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private String name;
    private String describe;
    private LocalDate date;

    public Task(){}
    
    public Task(String name, String describe){
        this.name = name;
        this.describe = describe;
    }

    public Long getId(){ return id; }

    public String getName(){ return name; }
    public void setName(String name){ this.name = name; }

    public String getDescribe(){ return describe; }
    public void setDescribe(String describe){ this.describe = describe;}

    public LocalDate getData(){ return date; }
    public void setData(LocalDate Date){ this.date = Date;}
}

