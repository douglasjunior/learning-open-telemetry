import Axios, { AxiosInstance } from 'axios';

import { AppConfigService } from '../../app-config/app-config.service';
import { Injectable } from '@nestjs/common';

type TodoResponse = {
  id: number;
  description: string;
}

export interface PlaceholderGateway {
  requestTodos(): Promise<TodoResponse[]>;
}

export const PlaceholderGateway = Symbol("PlaceholderGateway");

@Injectable()
export class PlaceholderGatewayHttp implements PlaceholderGateway {
  private readonly axios: AxiosInstance;

  constructor(appConfigService: AppConfigService) {
    this.axios = Axios.create({
      baseURL: appConfigService.placeholderEndpoint
    })
  }

  async requestTodos(): Promise<TodoResponse[]> {
    const response = await this.axios.get<TodoResponse[]>('/todos');
    return response.data;
  }
}
