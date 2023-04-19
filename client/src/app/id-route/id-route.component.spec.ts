import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IdRouteComponent } from './id-route.component';

describe('IdRouteComponent', () => {
  let component: IdRouteComponent;
  let fixture: ComponentFixture<IdRouteComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IdRouteComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(IdRouteComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
