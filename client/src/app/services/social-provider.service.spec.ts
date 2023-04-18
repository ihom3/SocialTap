import { TestBed } from '@angular/core/testing';

import { SocialProviderService } from './social-provider.service';

describe('SocialProviderService', () => {
  let service: SocialProviderService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SocialProviderService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
