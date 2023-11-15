import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ImagehistoryComponent } from './imagehistory.component';

describe('TagsComponent', () => {
  let component: ImagehistoryComponent;
  let fixture: ComponentFixture<ImagehistoryComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ImagehistoryComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ImagehistoryComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
