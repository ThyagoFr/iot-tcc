package github.thyagofr.api.domain.entity;

import com.sun.istack.NotNull;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;
import javax.validation.constraints.Email;
import javax.validation.constraints.NotBlank;
import javax.validation.constraints.Size;
import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;

@Entity
@Table(name = "users")
@EqualsAndHashCode(onlyExplicitlyIncluded = true)
@Getter
@Setter
@NoArgsConstructor
public class UserEntity implements Serializable {

  @GeneratedValue(strategy = GenerationType.IDENTITY)
  @Id
  @EqualsAndHashCode.Include
  @NotNull
  private Long id;

  @NotBlank
  @Size(min = 2, max = 255)
  private String name;

  @OneToMany(mappedBy = "user")
  private List<NotificationEntity> notifications = new ArrayList<>();

  @Email
  @NotBlank
  private String email;

  @NotBlank
  @Size(min = 8,max = 12)
  private String password;

}
