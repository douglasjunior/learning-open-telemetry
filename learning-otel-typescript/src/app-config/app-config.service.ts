import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';

@Injectable()
export class AppConfigService {

  constructor(private readonly configService: ConfigService) { }

  get databaseSynchronize(): boolean {
    return this.configService.get<boolean>('DATABASE_SYNCHRONIZE') ?? false;
  }

  get serverPort(): number {
    return this.configService.get<number>('PORT') ?? 3000;
  }

  get placeholderEndpoint(): string {
    return this.configService.getOrThrow<string>('PLACEHOLDER_ENDPOINT');
  }

}
