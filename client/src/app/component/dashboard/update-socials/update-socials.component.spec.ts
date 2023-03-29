import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UpdateSocialsComponent } from './update-socials.component';

describe('UpdateSocialsComponent', () => {
  let component: UpdateSocialsComponent;
  let fixture: ComponentFixture<UpdateSocialsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UpdateSocialsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(UpdateSocialsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
