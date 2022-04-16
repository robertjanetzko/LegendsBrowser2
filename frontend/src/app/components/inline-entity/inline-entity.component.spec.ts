import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InlineEntityComponent } from './inline-entity.component';

describe('InlineEntityComponent', () => {
  let component: InlineEntityComponent;
  let fixture: ComponentFixture<InlineEntityComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InlineEntityComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(InlineEntityComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
