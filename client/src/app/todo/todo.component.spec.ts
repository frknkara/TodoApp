import { ComponentFixture, fakeAsync, TestBed, tick } from '@angular/core/testing';
import { FormsModule } from '@angular/forms';
import { By } from '@angular/platform-browser';

import { TodoComponent } from './todo.component';

describe('TodoComponent', () => {
  let component: TodoComponent;
  let fixture: ComponentFixture<TodoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FormsModule],
      declarations: [TodoComponent]
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
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('h2')?.textContent).toBe('ToDo List');
  });

  it('should have a table', () => {
    let table = fixture.debugElement.query(By.css('table'));
    expect(table).toBeTruthy();
  });

  it('should have 2 columns', () => {
    let columns = fixture.debugElement.queryAll(By.css('table tr th'));
    let columnLength = columns.length;
    expect(columnLength).toBe(2);
  });

  it(`should have '#' column and 'Item' column`, () => {
    let columns = fixture.debugElement.queryAll(By.css('table tr th'));
    expect(columns[0].nativeElement.textContent).toBe('#');
    expect(columns[1].nativeElement.textContent).toBe('Item');
  });

  it('should have no record found message when list is empty', () => {
    const instance = fixture.componentInstance;
    spyOnProperty(instance, 'list').and.returnValue([]);
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('.no-record')?.textContent).toBe('No record found.');
  });

  it('should have an empty input text field', () => {
    let input = fixture.debugElement.query(By.css('input[type="text"]'));
    expect(input).toBeTruthy();
    expect(input.nativeElement.textContent).toBe('');
  });

  it('should have an add button', () => {
    let button = fixture.debugElement.query(By.css('button'));
    expect(button).toBeTruthy();
    expect(button.nativeElement.textContent).toBe('Add');
  });

  it('should check add button is disabled when text is empty', () => {
    const instance = fixture.componentInstance;
    spyOnProperty(instance, 'newItem').and.returnValue("");
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('#add-button')?.getAttribute('disabled')).toBe('');
  });

  it('should check add button is enabled when text is not empty', () => {
    const instance = fixture.componentInstance;
    spyOnProperty(instance, 'newItem').and.returnValue("some text");
    fixture.detectChanges();
    const compiled = fixture.nativeElement as HTMLElement;
    expect(compiled.querySelector('#add-button')?.getAttribute('disabled')).toBeNull();
  });

  it(`should bind input to 'newItem' property`, fakeAsync(() => {
    setInputValue('some text');

    const instance = fixture.componentInstance;
    expect(instance.newItem).toEqual('some text');
  }));

  function setInputValue(value: string) {
    let input = fixture.debugElement.query(By.css('#add-new-item-text')).nativeElement;
    input.value = value;
    input.dispatchEvent(new Event('input'));
    tick();
  }

  it(`should call 'onAddButtonClick' after button clicked`, fakeAsync(() => {
    setInputValue('some text');
    const instance = fixture.componentInstance;
    spyOn(instance, 'onAddButtonClick').and.callThrough();
    let button = fixture.debugElement.query(By.css('#add-button'));
    button.triggerEventHandler('click', null);
    tick();
    expect(instance.onAddButtonClick).toHaveBeenCalled();
  }));

  it('should add text to list when add button clicked', fakeAsync(() => {
    setInputValue('Some Item');
    const instance = fixture.componentInstance;
    spyOn(instance, 'onAddButtonClick').and.callThrough();
    let button = fixture.debugElement.query(By.css('#add-button'));
    button.triggerEventHandler('click', null);
    tick();
    expect(instance.list.length).toBeGreaterThan(0);
    let lastItem = instance.list[instance.list.length - 1];
    expect(lastItem.item).toEqual('Some Item');
  }));

  it('should clear input after add new item', fakeAsync(() => {
    setInputValue('Some Item');
    const instance = fixture.componentInstance;
    instance.onAddButtonClick();
    tick();
    expect(instance.newItem).toEqual('');
  }));
});
