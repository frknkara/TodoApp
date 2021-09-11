import { ComponentFixture, TestBed } from '@angular/core/testing';
import { By } from '@angular/platform-browser';

import { TodoComponent } from './todo.component';

describe('TodoComponent', () => {
  let component: TodoComponent;
  let fixture: ComponentFixture<TodoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TodoComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(TodoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it(`should have a heading as 'ToDo List'`, () => {
    const fixture = TestBed.createComponent(TodoComponent);
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('h2')?.textContent).toBe('ToDo List');
  });

  it('should have a table', () => {
    const fixture = TestBed.createComponent(TodoComponent);
    let table = fixture.debugElement.query(By.css('table'));
    expect(table).toBeTruthy();
  });

  it('should have 2 columns', () => {
    const fixture = TestBed.createComponent(TodoComponent);
    let columns = fixture.debugElement.queryAll(By.css('table tr th'));
    let columnLength = columns.length;
    expect(columnLength).toBe(2);
  });

  it(`should have '#' column and 'Item' column`, () => {
    const fixture = TestBed.createComponent(TodoComponent);
    let columns = fixture.debugElement.queryAll(By.css('table tr th'));
    expect(columns[0].nativeElement.textContent).toBe('#');
    expect(columns[1].nativeElement.textContent).toBe('Item');
  });
});
