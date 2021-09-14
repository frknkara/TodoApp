import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { environment } from 'src/environments/environment';
import { Observable } from 'rxjs';
import { TodoItem } from '../models/todo-item';

@Injectable({providedIn: 'root'})
export class TodoApiService {

  private _apiUrl = '';
  public get apiUrl() {
    return this._apiUrl;
  }
  public set apiUrl(value) {
    this._apiUrl = value;
  }

  constructor(private httpClient: HttpClient) { 
    this.apiUrl = environment.apiUrl;
  }
  
  private httpOptions = {
    headers: new HttpHeaders({
      'Content-Type': 'application/json'
    })
  }

  public getList() : Observable<any> {
    return this.httpClient.get(this.apiUrl + 'todo', this.httpOptions);
  }

  public add(item: TodoItem) : Observable<any> {
    return this.httpClient.post(this.apiUrl + 'todo', item, this.httpOptions);
  }
}