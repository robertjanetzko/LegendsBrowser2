import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CivilizationsComponent } from './civilizations.component';

describe('CivilizationsComponent', () => {
  let component: CivilizationsComponent;
  let fixture: ComponentFixture<CivilizationsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CivilizationsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CivilizationsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
