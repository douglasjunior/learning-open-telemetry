import { Body, Controller, Delete, Get, HttpCode, Param, Post } from '@nestjs/common';
import { TodoService } from './todo.service';
import { CreateTodoDto } from './dto/create-todo.dto';

@Controller('todo')
export class TodoController {

  constructor(private readonly todoService: TodoService) { }

  @Get('/error')
  @HttpCode(200)
  error() {
    return this.todoService.error();
  }

  @Get('/')
  get() {
    return this.todoService.getTodos();
  }

  @Get('/:todoId')
  getById(@Param('todoId') todoId: string) {
    return this.todoService.getTodoById(Number(todoId));
  }

  @Post('/')
  @HttpCode(201)
  create(@Body() createTodoDto: CreateTodoDto) {
    return this.todoService.createTodo(createTodoDto);
  }

  @Delete('/:todoId')
  @HttpCode(204)
  deleteById(@Param('todoId') todoId: string) {
    return this.todoService.deleteTodoById(Number(todoId));
  }

}
