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

  it('should have no record found message when list is empty', () => {
    const fixture = TestBed.createComponent(TodoComponent);
    const instance = fixture.componentInstance;
    spyOnProperty(instance, 'list').and.returnValue([]);
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('.no-record')?.textContent).toBe('No record found.');
  });

  it('should have an empty input text field', () => {
    const fixture = TestBed.createComponent(TodoComponent);
    let input = fixture.debugElement.query(By.css('input[type="text"]'));
    expect(input).toBeTruthy();
    expect(input.nativeElement.textContent).toBe('');
  });

  it('should have an add button', () => {
    const fixture = TestBed.createComponent(TodoComponent);
    let button = fixture.debugElement.query(By.css('button'));
    expect(button).toBeTruthy();
    expect(button.nativeElement.textContent).toBe('Add');
  });
});
