import { Component, OnInit } from '@angular/core';

interface TodoItem {
  id: number,
  item: string
}

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
  constructor() { }

  ngOnInit(): void {
  }

}
