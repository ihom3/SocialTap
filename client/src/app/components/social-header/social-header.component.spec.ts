import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SocialHeaderComponent } from './social-header.component';

describe('SocialHeaderComponent', () => {
  let component: SocialHeaderComponent;
  let fixture: ComponentFixture<SocialHeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SocialHeaderComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SocialHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
