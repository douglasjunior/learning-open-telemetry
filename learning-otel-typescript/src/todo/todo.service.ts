import { Inject, Injectable, Logger } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Repository } from 'typeorm';

import { Todo } from './data/todo.model';
import { CreateTodoDto } from './dto/create-todo.dto';
import { PlaceholderGateway } from './gateway/placeholder.gateway';

@Injectable()
export class TodoService {
  private readonly logger = new Logger(TodoService.name);

  constructor(
    @InjectRepository(Todo)
    private readonly todoRepository: Repository<Todo>,
    @Inject(PlaceholderGateway)
    private readonly placeholderGateway: PlaceholderGateway
  ) { }

  async getTodos(): Promise<object[]> {
    this.logger.log('Fetching all todos');
    const todoLists: object[][] = await Promise.all([
      this.todoRepository.find(),
      this.placeholderGateway.requestTodos()
    ]);
    return todoLists.flatMap(todoList => todoList);
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

  error() {
    this.logger.log('Create an error');
    const text: any = null;
    // eslint-disable-next-line @typescript-eslint/no-unsafe-call, @typescript-eslint/no-unsafe-member-access
    text.split();
  }
}
