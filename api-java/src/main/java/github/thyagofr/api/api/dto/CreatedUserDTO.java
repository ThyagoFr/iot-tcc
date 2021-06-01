package github.thyagofr.api.api.dto;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.io.Serializable;
import java.time.OffsetDateTime;

@Getter
@Setter
@NoArgsConstructor
public class CreatedUserDTO implements Serializable {
  private Long id;
  private OffsetDateTime createdAt;
  private String name;
  private String email;
  private String password;
}
