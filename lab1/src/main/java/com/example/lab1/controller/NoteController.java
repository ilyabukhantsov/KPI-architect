package com.example.lab1.controller;

import com.example.lab1.model.Note;
import com.example.lab1.repository.NoteRepository;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/notes")
public class NoteController {
    
    private final NoteRepository repository;

    public NoteController(NoteRepository repository) {
        this.repository = repository;
    }

    @GetMapping(produces = {MediaType.APPLICATION_JSON_VALUE, MediaType.TEXT_HTML_VALUE})
    public ResponseEntity<?> getAll(@RequestHeader(value = "Accept", defaultValue = "application/json") String accept) {
        List<Note> notes = repository.findAll();

        // Если запрос из браузера (Accept: text/html)
        if (accept.contains("text/html")) {
            StringBuilder html = new StringBuilder("<html><head><meta charset='UTF-8'></head><body>");
            html.append("<h1>Список заметок (Вариант 3858)</h1>");
            html.append("<table border='1' style='width:100%; border-collapse: collapse;'>");
            html.append("<tr style='background-color: #f2f2f2;'><th>ID</th><th>Заголовок</th><th>Действие</th></tr>");
            
            for (Note n : notes) {
                html.append("<tr>")
                    .append("<td>").append(n.getId()).append("</td>")
                    .append("<td>").append(n.getTitle()).append("</td>")
                    .append("<td><a href='/notes/").append(n.getId()).append("'>Открыть</a></td>")
                    .append("</tr>");
            }
            
            html.append("</table>");
            html.append("<br><p><i>Запрос обработан как text/html</i></p>");
            html.append("</body></html>");
            return ResponseEntity.ok().contentType(MediaType.TEXT_HTML).body(html.toString());
        }

        return ResponseEntity.ok(notes);
    }

    @PostMapping(consumes = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<Note> create(@RequestBody Note note) {
        Note savedNote = repository.save(note);
        return new ResponseEntity<>(savedNote, HttpStatus.CREATED);
    }

    @GetMapping(value = "/{id}", produces = {MediaType.APPLICATION_JSON_VALUE, MediaType.TEXT_HTML_VALUE})
    public ResponseEntity<?> getById(@PathVariable Long id, @RequestHeader(value = "Accept", defaultValue = "application/json") String accept) {
        Note note = repository.findById(id).orElse(null);
        
        if (note == null) {
            return ResponseEntity.status(HttpStatus.NOT_FOUND).body("Заметка не найдена");
        }

        if (accept.contains("text/html")) {
            String html = "<html><head><meta charset='UTF-8'></head><body>" +
                    "<h1>Просмотр заметки #" + note.getId() + "</h1>" +
                    "<p><b>Заголовок:</b> " + note.getTitle() + "</p>" +
                    "<p><b>Дата создания:</b> " + note.getCreatedAt() + "</p>" +
                    "<hr>" +
                    "<div style='background: #eee; padding: 10px;'>" + note.getContent() + "</div>" +
                    "<br><a href='/notes'>Вернуться к списку</a>" +
                    "</body></html>";
            return ResponseEntity.ok().contentType(MediaType.TEXT_HTML).body(html);
        }
        
        return ResponseEntity.ok(note);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> delete(@PathVariable Long id) {
        repository.deleteById(id);
        return ResponseEntity.noContent().build();
    }
}