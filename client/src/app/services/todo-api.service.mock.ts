import { of } from 'rxjs';
import { TodoItem } from '../models/todo-item';

export class TodoApiServiceStub {

  public getList() {
    return of([
      { id: 1, item: 'Çikolata' },
      { id: 2, item: 'Süt' }
    ]);
  }

  public add(item: TodoItem) {
    return of(item);
  }
}