package github.thyagofr.api.domain.repository;

import github.thyagofr.api.domain.entity.UserEntity;

import java.util.Optional;

public interface IUserRepository {

  UserEntity save(UserEntity user);

  Optional<UserEntity> findByID(Long userID);

  Optional<UserEntity> findByEmail(String email);

  void remove(Long userID);

}
