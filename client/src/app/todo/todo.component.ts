import { Component, OnInit } from '@angular/core';
import { TodoItem } from '../models/todo-item';
import { TodoApiService } from '../services/todo-api.service';

@Component({
  selector: 'app-todo',
  templateUrl: './todo.component.html',
  styleUrls: ['./todo.component.css']
})
export class TodoComponent implements OnInit {

  private _list: Array<TodoItem> = [];

  public get list() {
    return this._list;
  }
  public set list(value: Array<TodoItem>) {
    this._list = value;
  }

  private _newItem: string = "";

  public get newItem() {
    return this._newItem;
  }

  public set newItem(value: string) {
    this._newItem = value;
  }
  constructor(private todoApiService: TodoApiService) { }

  ngOnInit(): void {
    this.todoApiService.getList().subscribe((res) => {
      this.list = res;
    });
  }

  onAddButtonClick() {
    this.todoApiService.add({ id: undefined, item: this.newItem }).subscribe((res) => {
      this.list.push({ id: this.list.length, item: this.newItem });
      this.newItem = "";
    });
  }

}
