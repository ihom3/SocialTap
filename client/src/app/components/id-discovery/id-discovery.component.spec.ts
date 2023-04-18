import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IdDiscoveryComponent } from './id-discovery.component';

describe('IdDiscoveryComponent', () => {
  let component: IdDiscoveryComponent;
  let fixture: ComponentFixture<IdDiscoveryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ IdDiscoveryComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(IdDiscoveryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
