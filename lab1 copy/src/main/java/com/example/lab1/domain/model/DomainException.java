package com.example.lab1.domain.model;


public class DomainException extends RuntimeException {
    public DomainException(String message) {
        super(message);
    }
}