package github.thyagofr.api.api.assembler;

import github.thyagofr.api.api.dto.CreateUserDTO;
import github.thyagofr.api.api.dto.CreatedUserDTO;
import github.thyagofr.api.domain.entity.UserEntity;

import java.time.OffsetDateTime;

public final class UserAssembler {

  public static CreatedUserDTO toDTO(UserEntity entity) {
    CreatedUserDTO dto = new CreatedUserDTO();
    dto.setCreatedAt(OffsetDateTime.now());
    dto.setEmail(entity.getEmail());
    dto.setId(entity.getId());
    dto.setName(entity.getName());
    return dto;
  }

  public static UserEntity toEntity(CreateUserDTO dto) {
    UserEntity entity = new UserEntity();
    entity.setEmail(entity.getEmail());
    entity.setName(entity.getName());
    entity.setPassword(dto.getPassword());
    return entity;
  }
}
