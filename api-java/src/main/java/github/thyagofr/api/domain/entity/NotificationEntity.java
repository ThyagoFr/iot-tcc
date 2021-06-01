package github.thyagofr.api.domain.entity;

import com.sun.istack.NotNull;
import github.thyagofr.api.domain.enums.NotificationCondition;
import lombok.EqualsAndHashCode;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.*;
import javax.validation.constraints.NotBlank;
import java.io.Serializable;

@Entity
@Table(name = "notifications")
@EqualsAndHashCode(onlyExplicitlyIncluded = true)
@Getter
@Setter
@NoArgsConstructor
public class NotificationEntity implements Serializable {

  @GeneratedValue(strategy = GenerationType.IDENTITY)
  @Id
  @EqualsAndHashCode.Include
  @NotNull
  private Long id;

  private String deviceId;

  @ManyToOne
  private UserEntity user;

  private String parameter;

  @NotBlank
  @Enumerated(EnumType.STRING)
  private NotificationCondition condition;

  private Float value;

}
