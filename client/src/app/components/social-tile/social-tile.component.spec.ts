import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SocialTileComponent } from './social-tile.component';

describe('SocialTileComponent', () => {
  let component: SocialTileComponent;
  let fixture: ComponentFixture<SocialTileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SocialTileComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SocialTileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
