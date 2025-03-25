import { NestFactory } from '@nestjs/core';
import { ValidationPipe } from '@nestjs/common';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';

import { AppModule } from './app.module';
import { AppConfigService } from './app-config/app-config.service';
import { EntityNotFoundExceptionFilter } from './filters/entity-not-found-exception.filter';

const config = new DocumentBuilder()
  .setTitle('Todo example')
  .setVersion('1.0')
  .build();

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.useGlobalFilters(new EntityNotFoundExceptionFilter());
  app.useGlobalPipes(new ValidationPipe());
  const appConfigService = app.get(AppConfigService);

  const documentFactory = () => SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('swagger', app, documentFactory);

  await app.listen(appConfigService.serverPort);
  console.log('Server started on port: ' + appConfigService.serverPort)
}

bootstrap().catch(console.error);
