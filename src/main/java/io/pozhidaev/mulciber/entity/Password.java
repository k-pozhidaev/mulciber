package io.pozhidaev.mulciber.entity;

import lombok.Data;
import org.hibernate.annotations.GenericGenerator;
import org.hibernate.annotations.Immutable;

import javax.persistence.*;
import javax.validation.constraints.NotBlank;
import javax.validation.constraints.NotNull;
import javax.validation.constraints.Past;
import java.time.LocalDateTime;
import java.util.UUID;

@Data
@Entity
@Immutable
public class Password {

    @Id
    @GeneratedValue(generator = "UUID")
    @GenericGenerator(
            name = "UUID",
            strategy = "uuid"
    )
    @Column(name = "id", updatable = false, nullable = false, unique = true)
    private UUID id;

    @NotNull
    @NotBlank
    @Column(updatable = false, nullable = false)
    private String password;

    @Past
    @Column(name = "created_at")
    private LocalDateTime createdAt;

    @NotNull
    @ManyToOne(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
    @JoinColumn(name = "user_id", updatable = false)
    private User user;

    @PrePersist
    private void onCreate(){
        createdAt = LocalDateTime.now();
    }

}
