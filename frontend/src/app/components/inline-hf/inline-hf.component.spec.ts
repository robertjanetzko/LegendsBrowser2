import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InlineHfComponent } from './inline-hf.component';

describe('InlineHfComponent', () => {
  let component: InlineHfComponent;
  let fixture: ComponentFixture<InlineHfComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InlineHfComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(InlineHfComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
