package github.thyagofr.api.api.resource;

import github.thyagofr.api.api.assembler.UserAssembler;
import github.thyagofr.api.api.dto.CreateUserDTO;
import github.thyagofr.api.api.dto.CreatedUserDTO;
import github.thyagofr.api.domain.entity.UserEntity;
import github.thyagofr.api.domain.service.user.IUserService;
import lombok.AllArgsConstructor;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.*;

import javax.validation.Valid;

@RestController
@RequestMapping("/api/v1/users")
@AllArgsConstructor
public class User {

  private final IUserService userService;

  @PostMapping(
      produces = MediaType.APPLICATION_JSON_VALUE,
      consumes = MediaType.APPLICATION_JSON_VALUE
  )
  @ResponseStatus(HttpStatus.CREATED)
  public CreatedUserDTO register(@RequestBody @Valid CreateUserDTO body) {
    UserEntity user = this.userService.register(UserAssembler.toEntity(body));
    return UserAssembler.toDTO(user);
  }

}
