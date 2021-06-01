package github.thyagofr.api.domain.service.user;

import github.thyagofr.api.domain.entity.UserEntity;

public interface IUserService {
  UserEntity register(UserEntity user);
  UserEntity update(UserEntity user);
  UserEntity findByID(Long userID);
  UserEntity findByEmail(String email);
  void remove(Long userID);
}
