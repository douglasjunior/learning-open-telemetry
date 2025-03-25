import { Catch, ExceptionFilter, NotFoundException } from "@nestjs/common";
import { EntityNotFoundError } from 'typeorm/error/EntityNotFoundError'

@Catch(EntityNotFoundError)
export class EntityNotFoundExceptionFilter implements ExceptionFilter {
  public catch(exception: EntityNotFoundError) {
    throw new NotFoundException(exception.message);
  }
}
