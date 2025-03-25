import { Injectable, Logger } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';

import { Todo } from './data/todo.model';
import { CreateTodoDto } from './dto/create-todo.dto';

@Injectable()
export class TodoService {
  private readonly logger = new Logger(TodoService.name);

  constructor(
    @InjectRepository(Todo)
    private readonly todoRepository: Repository<Todo>,
  ) { }

  getTodos() {
    this.logger.log('Fetching all todos');
    return this.todoRepository.find();
  }

  getTodoById(id: number) {
    this.logger.log(`Fetching todo with id: ${id}`);
    return this.todoRepository.findOneByOrFail({ id })
  }

  async createTodo(createTodoDto: CreateTodoDto) {
    this.logger.log(`Creating todo with data: ${JSON.stringify(createTodoDto)}`);
    return this.todoRepository.save(createTodoDto);
  }

  deleteTodoById(id: number) {
    this.logger.log(`Deleting todo with id: ${id}`);
    return this.todoRepository.delete({ id });
  }

}
