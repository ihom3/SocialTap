import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SocialGridComponent } from './social-grid.component';

describe('SocialGridComponent', () => {
  let component: SocialGridComponent;
  let fixture: ComponentFixture<SocialGridComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SocialGridComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SocialGridComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
