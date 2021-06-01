package github.thyagofr.api.api.dto;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.validation.constraints.Email;
import javax.validation.constraints.NotBlank;
import javax.validation.constraints.Size;
import java.io.Serializable;

@Getter
@Setter
@NoArgsConstructor
public class CreateUserDTO implements Serializable {

  @NotBlank
  @Size(min = 2, max = 255, message = "a senha fornecida deve ter entre 2 e 255 caracteres")
  private String name;

  @Email(message = "informe um email válido")
  @NotBlank(message = "email obrigatório")
  private String email;

  @NotBlank
  @Size(min = 8, max = 12, message = "a senha fornecida deve ter entre 8 e 12 caracteres")
  private String password;

}
