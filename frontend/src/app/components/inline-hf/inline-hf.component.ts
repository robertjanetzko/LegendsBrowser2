import { Component, Input, OnInit } from '@angular/core';
import { EntityService } from 'src/app/entity.service';
import { HistoricalFigure } from 'src/app/types';

@Component({
  selector: 'app-inline-hf',
  templateUrl: './inline-hf.component.html',
  styleUrls: ['./inline-hf.component.scss']
})
export class InlineHfComponent implements OnInit {

  _id?: number
  data?: HistoricalFigure

  @Input()
  get id(): number | undefined {
    return this._id;
  }

  set id(val: number | undefined) {
    this._id = val
    if (val) {
      this.service.getOne<HistoricalFigure>("hf", val).then(data => this.data = data);
    }
  }

  constructor(
    private service: EntityService
  ) { }

  ngOnInit(): void {
  }

}
