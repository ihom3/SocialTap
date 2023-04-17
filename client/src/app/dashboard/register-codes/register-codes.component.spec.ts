import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterCodesComponent } from './register-codes.component';

describe('RegisterCodesComponent', () => {
  let component: RegisterCodesComponent;
  let fixture: ComponentFixture<RegisterCodesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RegisterCodesComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RegisterCodesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
