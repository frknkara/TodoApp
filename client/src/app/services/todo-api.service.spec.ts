import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { TodoApiService } from './todo-api.service';

const dummyTodoItemsResponse = {
  data : [
    { id: 1, item: 'Çikolata' },
    { id: 2, item: 'Süt' },
    { id: 3, item: 'Elma' },
    { id: 4, item: 'Ekmek' }
  ]
}

describe('TodoApiService', () => {
  let service: TodoApiService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule]
    });
    service = TestBed.inject(TodoApiService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  }); 

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it(`'getList' should return data`, () => {
    service.getList().subscribe((res) => {
      expect(res).toEqual(dummyTodoItemsResponse);
    });
    const req = httpMock.expectOne(service.apiUrl + 'todo');
    expect(req.request.method).toBe('GET');
    req.flush(dummyTodoItemsResponse);
  });

  it(`'add' should return success`, () => {
    service.add({id: 5, item: 'Şeker'}).subscribe((res) => {
      expect(res).toEqual({ msg: 'success' });
    });
    const req = httpMock.expectOne(service.apiUrl + 'todo');
    expect(req.request.method).toBe('POST');
    req.flush({ msg: 'success' });
  });
});