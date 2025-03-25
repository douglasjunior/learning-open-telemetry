import { Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { LoggerModule } from 'nestjs-pino';
import * as path from 'node:path';

import { TodoModule } from './todo/todo.module';
import { AppConfigModule } from './app-config/app-config.module';
import { AppConfigService } from './app-config/app-config.service';

@Module({
  imports: [
    LoggerModule.forRoot(),
    AppConfigModule,
    TypeOrmModule.forRootAsync({
      inject: [AppConfigService],
      useFactory: (appConfigService: AppConfigService) => {
        return {
          type: 'sqlite',
          database: path.join('database', 'db.sqlite'),
          entities: [path.join(__dirname, '**', '*.model.{ts,js}')],
          synchronize: appConfigService.databaseSynchronize,
        }
      },
    }),
    TodoModule,
  ],
  controllers: [],
  providers: [],
})
export class AppModule { }
