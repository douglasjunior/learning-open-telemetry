import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';

import { PlaceholderGateway, PlaceholderGatewayHttp } from './gateway/placeholder.gateway';
import { TodoController } from './todo.controller';
import { TodoService } from './todo.service';
import { Todo } from './data/todo.model';

@Module({
  imports: [TypeOrmModule.forFeature([Todo])],
  controllers: [TodoController],
  providers: [
    TodoService,
    {
      provide: PlaceholderGateway,
      useClass: PlaceholderGatewayHttp
    }
  ],
})
export class TodoModule { }
